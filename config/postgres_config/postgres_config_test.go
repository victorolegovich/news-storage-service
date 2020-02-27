package postgres_config

import (
	"go.uber.org/zap"
	"testing"
)

func TestNew(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
		return
	}

	if c := New(logger); c == nil {
		t.Error("it was not possible to get a configuration")
		return
	}
}

func TestConfig_String(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
		return
	}

	conf := New(logger)
	if conf == nil {
		t.Error("it was not possible to get a configuration")
		return
	}

	if conf.String() == "" {
		t.Error("couldn't get a line to connect to the database.")
		return
	}

}
