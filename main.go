package main

import (
	"github.com/nats-io/nats.go"
	config "github.com/victorolegovich/news-storage-service/config/nats_config"
	"github.com/victorolegovich/news-storage-service/storage"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		println(err.Error())
		return
	}

	conf := config.New(logger)

	broker, err := nats.Connect(conf.ServerURL)
	if err != nil {
		logger.Error(
			"failed to connect to the broker-server",
			zap.Error(err),
		)
	}
	logger.Info("the connection to the nats_config server was successful")

	store, err := storage.New()
	if err != nil {
		logger.Error("error during initialization of database connection", zap.Error(err))
		return
	}
	logger.Info("the connection to the database was successful")

	if _, err = getQueueSubscribe(broker, conf, newsHandler(logger, store)); err != nil {
		logger.Error("an error occurred when subscribing to the message queue.", zap.Error(err))
		return
	}

	if err = broker.Flush(); err != nil && broker.IsConnected() {
		logger.Error("message broker error", zap.Error(err))
		return
	}

	if err := broker.LastError(); err != nil {
		logger.Error("message broker error", zap.Error(err))
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	broker.Close()
	if err := store.Close(); err != nil {
		logger.Error("couldn't break up with the database.", zap.Error(err))
	}
	logger.Info("service has been stopped")
	os.Exit(1)
}
