package handlers

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"

	"github.com/temoon/telegram-bots/config"
)

type Redis interface {
	GetRedis() redis.UniversalClient
	ShutdownRedis() error
}

type RedisHandler struct {
	mu    sync.Mutex
	redis redis.UniversalClient
}

func (h *RedisHandler) GetRedis() redis.UniversalClient {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.redis != nil {
		return h.redis
	}

	h.redis = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: config.GetRedisAddress(),
	})

	return h.redis
}

func (h *RedisHandler) ShutdownRedis() (err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.redis != nil {
		if err = h.redis.Close(); err != nil {
			log.Error().Err(err).Msg("close redis")
		}
	}

	return
}
