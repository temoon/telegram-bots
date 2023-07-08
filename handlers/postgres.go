package handlers

import (
	"database/sql"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/temoon/telegram-bots/config"
)

type Postgres interface {
	GetPostgres() (*gorm.DB, error)
	ShutdownPostgres() error
}

type PostgresHandler struct {
	mu sync.Mutex
	db *gorm.DB
}

func (h *PostgresHandler) GetPostgres() (db *gorm.DB, err error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.db != nil {
		return h.db, nil
	}

	h.db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetPostgresUrl(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	return h.db, err
}

func (h *PostgresHandler) ShutdownPostgres() (err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.db != nil {
		var db *sql.DB
		if db, err = h.db.DB(); err != nil {
			return
		}

		return db.Close()
	}

	return
}
