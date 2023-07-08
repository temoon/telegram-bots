package handlers

import (
	"sync"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/rs/zerolog/log"

	"github.com/temoon/telegram-bots/config"
)

type ClickHouse interface {
	GetClickHouse() (clickhouse.Conn, error)
	ShutdownClickHouse() error
}

type ClickHouseHandler struct {
	mu sync.Mutex
	ch clickhouse.Conn
}

func (h *ClickHouseHandler) GetClickHouse() (ch clickhouse.Conn, err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.ch != nil {
		return h.ch, nil
	}

	h.ch, err = clickhouse.Open(config.GetClickHouseOpts())

	return h.ch, nil
}

func (h *ClickHouseHandler) ShutdownClickHouse() (err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.ch != nil {
		if err = h.ch.Close(); err != nil {
			log.Error().Err(err).Msg("close clickhouse")
		}
	}

	return
}
