package bots

import (
	"context"
	"sync"

	"github.com/temoon/telegram-bots-api"

	"github.com/temoon/telegram-bots/config"
)

type BaseHandler interface {
	OnUpdate(context.Context, *Request) error
	OnShutdown() error
	GetBot() *telegram.Bot
}

type Handler struct {
	mu  sync.Mutex
	bot *telegram.Bot
}

func (h *Handler) OnUpdate(ctx context.Context, req *Request) (err error) {
	return
}

func (h *Handler) OnShutdown() (err error) {
	return
}

func (h *Handler) GetBot() *telegram.Bot {
	h.Lock()
	defer h.Unlock()

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

func (h *Handler) Lock() {
	h.mu.Lock()
}

func (h *Handler) Unlock() {
	h.mu.Unlock()
}
