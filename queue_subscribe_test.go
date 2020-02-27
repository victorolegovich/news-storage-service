package main

import (
	"github.com/nats-io/nats.go"
	"github.com/victorolegovich/news-storage-service/config/nats_config"
	"github.com/victorolegovich/news-storage-service/storage"
	"go.uber.org/zap"
	"testing"
)

func TestQueueSubscribe(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err.Error())
		return
	}

	cfg := nats_config.New(logger)
	if cfg == nil {
		t.Error("configuration not obtained")
		return
	}

	conn, err := nats.Connect(cfg.ServerURL)
	if err != nil {
		t.Error(err.Error())
		return
	}

	store, err := storage.New()
	if err != nil {
		t.Error(err.Error())
	}

	handler := newsHandler(logger, store)

	if _, err = getQueueSubscribe(conn, cfg, handler); err != nil {
		t.Error(err.Error())
		return
	}
}
