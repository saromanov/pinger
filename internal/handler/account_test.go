package handler

import (
	"os"
	"testing"

	"github.com/saromanov/pinger/config"
	"github.com/saromanov/pinger/internal/models"
	"github.com/saromanov/pinger/internal/storage"
	"github.com/stretchr/testify/assert"
)

var hand *Handler

func init() {
	st, err := storage.New(&config.Config{
		Name:     os.Getenv("PINGER_TEST_DB_NAME"),
		Password: os.Getenv("PINGER_TEST_DB_PASSWORD"),
		User:     os.Getenv("PINGER_TEST_DB_USER"),
	})
	if err != nil {
		panic(err)
	}

	handTmp, err := New(st)
	if err != nil {
		panic(err)
	}

	hand = handTmp
}

func TestCreateAccount(t *testing.T) {
	m := &models.Account{}
	_, _, err := hand.CreateAccount(m)
	assert.Error(t, err, "should return error")
}
