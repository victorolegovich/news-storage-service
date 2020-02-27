package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/victorolegovich/news-storage-service/config/nats_config"
	"github.com/victorolegovich/news-storage-service/storage"
	"go.uber.org/zap"
)

func getQueueSubscribe(c *nats.Conn, cfg *nats_config.Config, handler func(msg *nats.Msg)) (*nats.Subscription, error) {
	return c.QueueSubscribe(cfg.Subject, cfg.NewsQueue, handler)
}

func newsHandler(logger *zap.Logger, store *storage.Storage) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		logger.Info("New request received.", zap.String("message", string(msg.Data)))

		ID := string(msg.Data)

		newsItem, err := store.GetNewsItemByID(ID)
		if err != nil {
			logger.Error("no news with this ID was found.", zap.Error(err))
			if err = msg.Respond([]byte("no record was found")); err != nil {
				logger.Error(
					"failed to respond to a message from the sender of the request",
					zap.String("sub", msg.Subject),
					zap.Error(err),
				)
			}
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
	}
}
