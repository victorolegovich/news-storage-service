package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/victorolegovich/news-storage-service/proto"
)

func New() (*Storage, error) {
	conn, err := pgx.Connect(context.Background(), connAddr())
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

func (p *Storage) GetNewsItemByID(ID string) (item proto.NewsItem, err error) {
	item.ID = ID

	err = p.conn.QueryRow(context.Background(),
		"SELECT header FROM public.news WHERE id = $1", ID).
		Scan(&item.Header)

	return
}

func connAddr() string {
	const (
		h    = "localhost"
		p    = "5032"
		u    = "postgres"
		pass = "Dbrnjh777"
		dbn  = "postgres"
		ssl  = "disable"
	)

	return "h=" + h + " p=" + p + " u=" + u + " pass=" + pass + " dbn=" + dbn + " sslmode=" + ssl
}

func createTableIfNotExists(c *pgx.Conn) error {
	sql := `CREATE TABLE IF NOT EXISTS public.news
(
    id text NOT NULL,
    header text NOT NULL,
    PRIMARY KEY (id)
);`

	_, err := c.Exec(context.Background(), sql)
	return err
}
