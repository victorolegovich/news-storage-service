package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/victorolegovich/news-storage-service/storage"
	"go.uber.org/zap"
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

	if _, err = broker.QueueSubscribe("client", "news", reply(logger)); err != nil {
		logger.Error("an error occurred when subscribing to the message queue.", zap.Error(err))
	}

	if err = broker.Flush(); err != nil {
		logger.Error("message broker error", zap.Error(err))
	}

	if err := broker.LastError(); err != nil {
		logger.Error("message broker error", zap.Error(err))
	}

	if err != nil {
		logger.Error(
			"failed to sign up for the news-request queue",
			zap.Error(err),
		)
	}

}

func reply(logger *zap.Logger) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		store := storage.New()

		ID := string(msg.Data)

		newsItem, err := store.GetNewsItemByID(ID)
		if err != nil {
			logger.Error("no news with this ID was found.", zap.Error(err))
		}

		message, err := proto.Marshal(newsItem)
		if err != nil {
			logger.Error("message could not be converted", zap.Error(err))
		}

		err = msg.Respond(message)
		if err != nil {
			logger.Error(
				"message could not be answered",
				zap.Error(err),
				zap.String("sub", msg.Subject),
			)
		}
	}
}
