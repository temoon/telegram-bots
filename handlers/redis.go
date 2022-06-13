package handlers

import (
	"sync"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"

	"github.com/temoon/telegram-bots/config"
)

type Redis interface {
	GetRedis() redis.UniversalClient
	Shutdown() error
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

func (h *RedisHandler) Shutdown() (err error) {
	if h.redis != nil {
		if err = h.redis.Close(); err != nil {
			log.WithError(err).Error("close redis")
		}
	}

	return
}
