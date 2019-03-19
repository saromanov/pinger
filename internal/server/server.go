// Package server defines rest api endpoints for pinger
package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saromanov/pinger/internal/handler"
	"github.com/saromanov/pinger/internal/models"
)

type server struct {
	hand   *handler.Handler
	router *mux.Router
}

// createAccount makes a new account
func (s *server) createAccount() http.HandleFunc {
	account := &models.Account{}
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
		if err != nil {
			return
		}
		resp := account.Create()
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
