package storage

import (
	"context"
	"github.com/jackc/pgx"
	"github.com/victorolegovich/news-storage-service/proto"
)

func New() *Storage {
	conn, err := pgx.Connect(context.Background(), "")
	if err != nil {
		println(err.Error())
		return nil
	}

	return &Storage{connection: conn}
}

type Storage struct {
	connection *pgx.Conn
}

func (p *Storage) GetNewsItemByID(ID string) (ni *proto.NewsItem, err error) {
	return
}
