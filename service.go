package main

import (
	"context"
	"github.com/victorolegovich/news-storage-service/proto"
)

type newsStoreService struct {
}

func newService() *newsStoreService {
	return nil
}

func (service newsStoreService) GetNewsItem(ctx context.Context, req *proto.NewsRequest) (item *proto.NewsItem, err error) {
	return
}
