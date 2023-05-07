package handlers

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/temoon/telegram-bots/config"
)

type Postgres interface {
	GetPostgres()
	Shutdown() error
}

type PostgresHandler struct {
	mu     sync.Mutex
	pgPool *pgxpool.Pool
}

func (h *PostgresHandler) GetPostgresPool(ctx context.Context) (postgres *pgxpool.Pool, err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.pgPool != nil {
		return h.pgPool, nil
	}

	if h.pgPool, err = pgxpool.New(ctx, config.GetPostgresUrl()); err != nil {
		return
	}

	return h.pgPool, nil
}

func (h *PostgresHandler) Shutdown() (err error) {
	if h.pgPool != nil {
		h.pgPool.Close()
	}

	return
}
