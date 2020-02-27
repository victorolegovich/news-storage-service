package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/victorolegovich/news-storage-service/config/postgres_config"
	"go.uber.org/zap"
	"testing"
)

func TestCreateTableIfNotExists(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err.Error())
		return
	}

	connString := postgres_config.NewConfig(logger).String()
	if connString == "" {
		t.Error("blank connection string")
		return
	}

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Error(err)
		return
	}

	if err = createTableIfNotExists(conn); err != nil {
		t.Error(err.Error())
		return
	}
}

func TestNew(t *testing.T) {
	if _, err := New(); err != nil {
		t.Error(err.Error())
		return
	}
}

func TestStorage_GetNewsItemByID(t *testing.T) {
	store, err := New()
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = store.GetNewsItemByID("xklkasqfqcxx")
	if err != pgx.ErrNoRows{
		t.Error(err)
		return
	}
}
