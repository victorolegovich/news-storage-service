package storage

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
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

func (p *Storage) GetNewsItemByID(ID int32) (ni *proto.NewsItem) {
	var (
		header string
		date   *timestamp.Timestamp
	)

	err := p.connection.QueryRow(context.Background(), "", ID).Scan(header, date)
	if err != nil {
		println(err.Error())
		return
	}
	return
}
