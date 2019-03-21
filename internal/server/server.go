// Package server defines rest api endpoints for pinger
package server

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/saromanov/pinger/config"
	"github.com/saromanov/pinger/internal/handler"
	"github.com/saromanov/pinger/internal/log"
	"github.com/saromanov/pinger/internal/models"
	pb "github.com/saromanov/pinger/proto"
)

type server struct {
	hand    *handler.Handler
	router  *mux.Router
	address string
}

// createAccount makes a new account
func (s *server) createAccount(w http.ResponseWriter, r *http.Request) {
	account := &pb.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		return
	}
	_, err = s.hand.CreateAccount(&models.Account{
		Email:    account.Email,
		Password: account.Password,
		Name:     account.Name,
	})
	if err != nil {
		return
	}
}

func (s *server) makeHandlers() {
	s.router.HandleFunc("/v1/users", s.createAccount)
}

func (s *server) startServer() {
	log.Infof("server is started at %s")
	srv := &http.Server{
		Addr:         s.address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Info("shutting down server")
}

// New makes http endpoints and handler
func New(h *handler.Handler, c *config.Config) {
	r := mux.NewRouter()
	s := &server{
		hand:   h,
		router: r,
	}
	s.makeHandlers()
	s.startServer()
}
