// Package server defines rest api endpoints for pinger
package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saromanov/pinger/internal/handler"
	"github.com/saromanov/pinger/internal/models"
	"github.com/saromanov/pinger/proto"
)

type server struct {
	hand   *handler.Handler
	router *mux.Router
}

// createAccount makes a new account
func (s *server) createAccount() http.HandleFunc {
	account := &proto.Account{}
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(account)
		if err != nil {
			return
		}
		_, err = s.hand.CreateAccount(&models.Account{
			Email: account.Email,
			Password: account.Password,
			Name: account.Name,
		})
		if err != nil {
			return
		}
	}
}

func (s *server) makeHandlers() {
	s.router.HandleFunc("/v1/users", s.createAccount)
}

// Create makes http endpoints and handler
func Create(h *handler.Handler) {
	r := mux.NewRouter()
	s := &server{
		hand:   h,
		router: r,
	}
	s.makeHandlers()
}
