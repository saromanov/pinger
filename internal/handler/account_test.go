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

func TestCreateFailedAccount(t *testing.T) {
	_, _, err := hand.CreateAccount(&models.Account{})
	assert.EqualError(t, err, errNoEmail.Error(), "should return error")

	_, _, err = hand.CreateAccount(&models.Account{Email: "123@mail.ru"})
	assert.EqualError(t, err, errNoPassword.Error(), "should return error")

	_, _, err = hand.CreateAccount(&models.Account{Email: "mail.ru"})
	assert.Error(t, err, "should return error")

	_, _, err = hand.CreateAccount(&models.Account{Password: "123"})
	assert.EqualError(t, err, errNoEmail.Error(), "should return error")

}
