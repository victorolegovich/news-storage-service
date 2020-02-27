package nats_config

import (
	"go.uber.org/zap"
	"testing"
)

func TestNew(t *testing.T) {
	logger, err :=  zap.NewDevelopment()
	if err != nil {
		t.Error(err)
		return
	}

	if c := New(logger); c ==nil{
		t.Error("it was not possible to get a configuration")
		return
	}
}
