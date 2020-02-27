package nats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"os"
)

const (
	envNatshost  = "NATS_HOST"
	envNatssub   = "NATS_SUB"
	envNewsQueue = "NATS_NEWS_QUEUE"

	defaultNatshost  = nats.DefaultURL
	defaultNatssub   = "storage"
	defaultNewsQueue = "news"
)

type Config struct {
	ServerURL, Subject, NewsQueue string
}

func New(logger *zap.Logger) *Config {
	url := os.Getenv(envNatshost)
	if url == "" {
		logger.Error(
			"the environment variable was not set with NATS_HOST - it is set by default",
			zap.String("NATS_HOST", "can cause an error when connecting to the nats server"),
		)
		url = defaultNatshost
	}

	sub := os.Getenv(envNatssub)
	if sub == "" {
		logger.Error(
			"the environment variable was not set with NATS_SUB - it is set by default",
			zap.String("NATS_SUB", "can cause a customer to make a mistake"),
		)
		sub = defaultNatssub
	}

	queue := os.Getenv(envNewsQueue)
	if queue == "" {
		logger.Error(
			"the environment variable was not set with NATS_NEWS_QUEUE - it is set by default",
			zap.String("NATS_NEWS_QUEUE", "can cause a customer to make a mistake"),
		)
		queue = defaultNewsQueue
	}

	return &Config{ServerURL: url, Subject: sub, NewsQueue: queue}
}
