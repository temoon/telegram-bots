package handlers

import (
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/temoon/telegram-bots/config"
)

type Kafka interface {
	GetKafka(string) (*kgo.Client, error)
	ShutdownKafka() error
}

type KafkaHandler struct {
	mu    sync.Mutex
	kafka *kgo.Client
}

func (h *KafkaHandler) GetKafka(group string) (kafka *kgo.Client, err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.kafka != nil {
		return h.kafka, nil
	}

	opts := []kgo.Opt{
		kgo.SeedBrokers(config.GetKafkaSeeds()...),
	}

	if group != "" {
		opts = append(opts, kgo.ConsumerGroup(group))
	}

	if h.kafka, err = kgo.NewClient(opts...); err != nil {
		return
	}

	return h.kafka, nil
}

func (h *KafkaHandler) ShutdownKafka() (err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.kafka != nil {
		h.kafka.Close()
	}

	return
}
