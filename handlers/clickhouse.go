package handlers

import (
	"sync"

	"github.com/ClickHouse/clickhouse-go/v2"
	log "github.com/sirupsen/logrus"

	"github.com/temoon/telegram-bots/config"
)

type ClickHouse interface {
	GetClickHouse()
	Shutdown() error
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

	opts := config.GetClickHouseOpts()
	if h.ch, err = clickhouse.Open(opts); err != nil {
		return
	}

	return h.ch, nil
}

func (h *ClickHouseHandler) Shutdown() (err error) {
	if h.ch != nil {
		if err = h.ch.Close(); err != nil {
			log.WithError(err).Error("close clickhouse")
		}
	}

	return
}
