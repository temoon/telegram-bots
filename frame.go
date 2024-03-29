package bots

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/temoon/telegram-bots-api"
	"github.com/temoon/telegram-bots-api/requests"

	"github.com/temoon/telegram-bots/config"
	. "github.com/temoon/telegram-bots/helpers"
)

type State int

type Frame struct {
	state  State
	mu     sync.Mutex
	server *http.Server

	Handler Handler
	BotUser *telegram.User
}

const (
	StateIdle State = iota
	StateRunning
	StateShutdown
)

func (f *Frame) Run(ctx context.Context) (err error) {
	if f.BotUser == nil {
		req := requests.GetMe{}

		var res interface{}
		if res, err = req.Call(ctx, f.Handler.GetBot()); err != nil {
			return
		}

		var ok bool
		if f.BotUser, ok = res.(*telegram.User); !ok {
			return errors.New("error while getting bot user")
		}
	}

	if config.IsHttpServer() {
		return f.Listen(ctx)
	}

	return f.Loop(ctx)
}

func (f *Frame) Loop(ctx context.Context) (err error) {
	if err = f.onRun(); err != nil {
		return
	}

	req := requests.GetUpdates{
		Timeout: Int64(1),

		AllowedUpdates: []string{
			telegram.UpdatesAllowedMessage,
			telegram.UpdatesAllowedCallbackQuery,
		},
	}

	var res interface{}
	for f.isRunning() {
		if res, err = req.Call(ctx, f.Handler.GetBot()); err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}

			log.Error().Err(err).Msg("get updates")
			time.Sleep(time.Second)

			continue
		}

		for _, update := range *res.(*[]telegram.Update) {
			req.Offset = Int64(update.UpdateId + 1)

			if err = f.onUpdate(ctx, &update); err != nil {
				log.Error().Err(err).Msg("on update")
				continue
			}
		}
	}

	return
}

func (f *Frame) Listen(ctx context.Context) (err error) {
	if err = f.onRun(); err != nil {
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc(config.GetHttpEndpoint(), func(res http.ResponseWriter, req *http.Request) {
		if !f.isRunning() {
			log.Error().Err(err).Msg("not running")
			res.WriteHeader(http.StatusNotFound)
			return
		}

		var err error
		var update telegram.Update
		if err = json.NewDecoder(req.Body).Decode(&update); err != nil {
			log.Error().Err(err).Msg("error while decoding json payload")
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = f.onUpdate(ctx, &update); err != nil {
			log.Error().Err(err).Msg("on update")

			var errWithStatusCode *ErrorWithStatusCode
			if errors.As(err, &errWithStatusCode) {
				res.WriteHeader(errWithStatusCode.StatusCode)
			}

			return
		}

		res.WriteHeader(http.StatusOK)
	})

	f.server = &http.Server{
		Addr:         config.GetHttpAddress(),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	certFile := config.GetHttpCertFile()
	certKey := config.GetHttpCertKey()
	if certFile != "" && certKey != "" {
		var pemCerts []byte
		if pemCerts, err = os.ReadFile(certFile); err != nil {
			return
		}

		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(pemCerts)

		f.server.TLSConfig = &tls.Config{
			ServerName: config.GetHttpServerName(),
			ClientAuth: tls.RequestClientCert,
			ClientCAs:  certPool,
			MinVersion: tls.VersionTLS12,
		}

		log.Debug().Msg("listening for updates in secure mode...")
		if err = f.server.ListenAndServeTLS(certFile, certKey); errors.Is(err, http.ErrServerClosed) {
			return nil
		} else {
			return
		}
	} else {
		log.Debug().Msg("listening for updates...")
		if err = f.server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			return nil
		} else {
			return
		}
	}
}

func (f *Frame) OnInterrupt(cancel context.CancelFunc) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt)

	for !f.isShuttingDown() {
		sig := <-osSignal
		log.Debug().Str("sig", sig.String()).Msg("signal received")

		if sig == os.Interrupt {
			f.Shutdown(cancel)
			break
		}
	}
}

func (f *Frame) Shutdown(cancel context.CancelFunc) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.state == StateShutdown {
		return
	}

	f.state = StateShutdown
	cancel()

	var err error
	if err = f.Handler.OnShutdown(); err != nil {
		log.Error().Err(err).Msg("error on shutting down handler")
	}

	if f.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err = f.server.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("error while shutting down the HTTP server")
		}
	}
}

func (f *Frame) onRun() (err error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.state != StateIdle {
		return errors.New("already running")
	}

	f.state = StateRunning

	return
}

func (f *Frame) onUpdate(ctx context.Context, u *telegram.Update) (err error) {
	log.Debug().Interface("update", u).Msg("update received")

	var req *Request
	switch {
	case u.CallbackQuery != nil:
		if u.CallbackQuery.Message == nil || u.CallbackQuery.From.IsBot {
			return &ErrorWithStatusCode{
				StatusCode: http.StatusBadRequest,
				Err:        errors.New("invalid callback query"),
			}
		}

		if req, err = ParseCallbackQuery(u.CallbackQuery); err != nil {
			return
		}
	case u.Message != nil:
		if u.Message.From == nil || u.Message.From.IsBot {
			return &ErrorWithStatusCode{
				StatusCode: http.StatusBadRequest,
				Err:        errors.New("invalid message"),
			}
		}

		if req, err = ParseMessage(u.Message); err != nil {
			return
		}
	default:
		return &ErrorWithStatusCode{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("unexpected update"),
		}
	}

	if f.BotUser != nil {
		req.BotUsername = StringOrEmpty(f.BotUser.Username)
	}

	if !config.IsBotUserAllowed(req.UserId) {
		return &ErrorWithStatusCode{
			StatusCode: http.StatusOK,
			Err:        errors.New("user not allowed"),
		}
	}

	if err = f.Handler.OnUpdate(ctx, req); err != nil {
		SendErrorMessage(ctx, f.Handler.GetBot(), req.ChatId.GetId())
	}

	return
}

func (f *Frame) isRunning() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.state == StateRunning
}

func (f *Frame) isShuttingDown() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.state == StateShutdown
}
