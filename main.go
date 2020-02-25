package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
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

	broker, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logger.Error(
			"failed to connect to the broker-server",
			zap.Error(err),
		)
	}
	logger.Info("the connection to the nats server was successful")

	store, err := storage.New()
	if err != nil {
		logger.Error("error during initialization of database connection", zap.Error(err))
		return
	}
	logger.Info("the connection to the database was successful")

	if _, err = broker.QueueSubscribe("storage", "news", func(msg *nats.Msg) {
		logger.Info("New request received.", zap.String("message", string(msg.Data)))

		ID := string(msg.Data)

		newsItem, err := store.GetNewsItemByID(ID)
		if err != nil {
			logger.Error("no news with this ID was found.", zap.Error(err))
			msg.Respond([]byte("no record was found"))
			return
		}

		message, err := proto.Marshal(&newsItem)
		if err != nil {
			logger.Error("message could not be converted", zap.Error(err))
			return
		}

		err = msg.Respond(message)
		if err != nil {
			logger.Error(
				"message could not be answered",
				zap.Error(err),
				zap.String("sub", msg.Subject),
			)
			return
		}
	}); err != nil && broker.IsConnected() {
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
