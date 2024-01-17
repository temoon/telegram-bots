package handlers

import (
	"context"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"sync"

	"github.com/temoon/telegram-bots/config"
)

type YandexDB interface {
	GetYandexDB() (*ydb.Driver, error)
	ShutdownYandexDB() error
}

type YandexDBHandler struct {
	mu sync.Mutex
	db *ydb.Driver
}

func (h *YandexDBHandler) GetYandexDB(ctx context.Context) (db *ydb.Driver, err error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.db != nil {
		return h.db, nil
	}

	h.db, err = ydb.Open(ctx, config.GetYdbConnectionString(), config.GetYdbOpts()...)

	return h.db, err
}

func (h *YandexDBHandler) ShutdownYandexDB(ctx context.Context) (err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.db != nil {
		return h.db.Close(ctx)
	}

	return
}
