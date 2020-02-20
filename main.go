package main

import (
	"context"
	proto "github.com/victorolegovich/news-storage-service/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

const port string = ":50051"


func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		println(err)
		return
	}

	defer logger.Sync()

	_, err = net.Listen("tcp", port)
	if err != nil {
		logger.Info("can't listen to the port :50051",
			zap.Error(err),
		)
	}

	logger.Info("listening port :50051")


}
