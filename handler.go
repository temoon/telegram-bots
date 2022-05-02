package bots

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/temoon/telegram-bots-api"
	. "github.com/temoon/telegram-bots-api/helpers"
	"github.com/temoon/telegram-bots-api/requests"

	"github.com/temoon/telegram-bots/config"
	"github.com/temoon/telegram-bots/helpers"
)

type State int

type Handler struct {
	state  State
	mu     sync.Mutex
	server *http.Server

	bot *telegram.Bot
}

type Router func(context.Context, *Handler, *Update) error

const (
	StateIdle State = iota
	StateRunning
	StateShutdown
)

func (h *Handler) Run(ctx context.Context, router Router) (err error) {
	if config.IsHttpServer() {
		return h.Listen(ctx, router)
	}

	return h.Loop(ctx, router)
}

func (h *Handler) Loop(ctx context.Context, router Router) (err error) {
	if err = h.onRun(); err != nil {
		return
	}

	req := requests.GetUpdates{
		Timeout: Int32(1),

		AllowedUpdates: []string{
			telegram.UpdatesAllowedMessage,
			telegram.UpdatesAllowedCallbackQuery,
		},
	}

	var res interface{}
	for h.isRunning() {
		if res, err = req.Call(ctx, h.GetBot()); err != nil {
			if urlErr, ok := err.(*url.Error); ok && urlErr.Err == context.Canceled {
				return
			}

			log.WithError(err).Error("get updates")
			time.Sleep(time.Second)

			continue
		}

		for _, update := range *res.(*[]telegram.Update) {
			req.Offset = Int32(update.UpdateId + 1)

			if err = h.onUpdate(ctx, &update, router); err != nil {
				log.WithError(err).Error("on update")
				continue
			}
		}
	}

	return
}

func (h *Handler) Listen(ctx context.Context, router Router) (err error) {
	if err = h.onRun(); err != nil {
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc(config.GetHttpEndpoint(), func(res http.ResponseWriter, req *http.Request) {
		if !h.isRunning() {
			log.WithError(err).Error("not running")
			res.WriteHeader(http.StatusNotFound)
			return
		}

		var err error
		var update telegram.Update
		if err = json.NewDecoder(req.Body).Decode(&update); err != nil {
			log.WithError(err).Error("error while decoding json payload")
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = h.onUpdate(ctx, &update, router); err != nil {
			log.WithError(err).Error("on update")

			var errWithStatusCode *ErrorWithStatusCode
			if errors.As(err, &errWithStatusCode) {
				res.WriteHeader(errWithStatusCode.StatusCode)
			} else {
				res.WriteHeader(http.StatusInternalServerError)
			}

			return
		}

		res.WriteHeader(http.StatusOK)
	})

	h.server = &http.Server{
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
		if pemCerts, err = ioutil.ReadFile(certFile); err != nil {
			return
		}

		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(pemCerts)

		h.server.TLSConfig = &tls.Config{
			ServerName: config.GetHttpServerName(),
			ClientAuth: tls.RequestClientCert,
			ClientCAs:  certPool,
			MinVersion: tls.VersionTLS12,
		}

		log.Debug("listening for updates in secure mode...")
		if err = h.server.ListenAndServeTLS(certFile, certKey); errors.Is(err, http.ErrServerClosed) {
			return nil
		} else {
			return
		}
	} else {
		log.Debug("listening for updates...")
		if err = h.server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
			return nil
		} else {
			return
		}
	}
}

func (h *Handler) GetBot() *telegram.Bot {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.bot != nil {
		return h.bot
	}

	h.bot = telegram.NewBot(&telegram.BotOpts{
		Token:   config.GetBotToken(),
		Timeout: config.GetBotTimeout(),
		Env:     config.GetBotEnvironment(),
	})

	return h.bot
}

func (h *Handler) OnInterrupt(cancel context.CancelFunc) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt)

	for !h.isShuttingDown() {
		sig := <-osSignal
		log.WithField("sig", sig.String()).Debug("signal received")

		if sig == os.Interrupt {
			h.Shutdown(cancel)
			break
		}
	}
}

func (h *Handler) Shutdown(cancel context.CancelFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.state == StateShutdown {
		return
	}

	h.state = StateShutdown
	cancel()

	if h.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := h.server.Shutdown(ctx); err != nil {
			log.WithError(err).Error("error while shutting down the HTTP server")
		}
	}
}

func (h *Handler) onRun() (err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.state != StateIdle {
		return errors.New("already running")
	}

	h.state = StateRunning

	return
}

func (h *Handler) onUpdate(ctx context.Context, u *telegram.Update, router Router) (err error) {
	log.WithField("update", u).Debug("update received")

	var update *Update
	switch {
	case u.CallbackQuery != nil:
		if u.CallbackQuery.Message == nil || u.CallbackQuery.From.IsBot {
			return &ErrorWithStatusCode{
				StatusCode: http.StatusBadRequest,
				Err:        errors.New("invalid callback query"),
			}
		}

		if update, err = ParseCallbackQuery(u.CallbackQuery); err != nil {
			return
		}
	case u.Message != nil:
		if u.Message.From == nil || u.Message.From.IsBot {
			return &ErrorWithStatusCode{
				StatusCode: http.StatusBadRequest,
				Err:        errors.New("invalid message"),
			}
		}

		if update, err = ParseMessage(u.Message); err != nil {
			return
		}
	default:
		return &ErrorWithStatusCode{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("unexpected update"),
		}
	}

	if !config.IsBotUserAllowed(update.UserId) {
		return &ErrorWithStatusCode{
			StatusCode: http.StatusOK,
			Err:        errors.New("user not allowed"),
		}
	}

	if err = router(ctx, h, update); err != nil {
		helpers.SendErrorMessage(ctx, h.GetBot(), update.ChatId)
	}

	return
}

func (h *Handler) isRunning() bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	return h.state == StateRunning
}

func (h *Handler) isShuttingDown() bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	return h.state == StateShutdown
}
