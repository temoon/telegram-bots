package bots

import (
	"context"
	"sync"

	"github.com/temoon/telegram-bots-api"

	"github.com/temoon/telegram-bots/config"
)

type Handler interface {
	OnUpdate(context.Context, *Request) error
	OnShutdown() error
	GetBot() *telegram.Bot
}

type BaseHandler struct {
	mu  sync.Mutex
	bot *telegram.Bot
}

//goland:noinspection GoUnusedParameter
func (h *BaseHandler) OnUpdate(ctx context.Context, req *Request) (err error) {
	return
}

func (h *BaseHandler) OnShutdown() (err error) {
	return
}

func (h *BaseHandler) GetBot() *telegram.Bot {
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
