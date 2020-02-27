package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/victorolegovich/news-storage-service/config/postgres_config"
	"github.com/victorolegovich/news-storage-service/proto"
	"go.uber.org/zap"
)

func New() (*Storage, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		println("failed to create a logger zap")
		return nil, err
	}
	connString := postgres_config.New(logger).String()
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	if err = createTableIfNotExists(conn); err != nil {
		return nil, err
	}

	return &Storage{conn: conn}, nil
}

type Storage struct {
	conn *pgx.Conn
}

func (s *Storage) GetNewsItemByID(ID string) (item proto.NewsItem, err error) {
	item.ID = ID

	println(ID)

	err = s.conn.QueryRow(context.Background(),
		"select header,to_char(creation_date,'DD-MM-YYYY') from public.news where id = $1", ID).
		Scan(&item.Header, &item.CreationDate)

	return
}

func (s *Storage) Close() error {
	return s.conn.Close(context.Background())
}

func createTableIfNotExists(c *pgx.Conn) error {
	sql := `CREATE TABLE IF NOT EXISTS public.news
(
    id text NOT NULL,
    header text NOT NULL,
	creation_date date,
    PRIMARY KEY (id)
);`

	_, err := c.Exec(context.Background(), sql)
	return err
}
