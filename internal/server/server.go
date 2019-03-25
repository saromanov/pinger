// Package server defines rest api endpoints for pinger
package server

import (
	"context"
	"encoding/json"
	"fmt"
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
	if account.Email == "" && account.Password == "" {
		http.Error(w, "email or password is not defined", http.StatusBadRequest)
		return
	}
	_, err = s.hand.CreateAccount(&models.Account{
		Email:    account.Email,
		Password: account.Password,
		Name:     account.Name,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to create account: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// createSite makes a new site for user
func (s *server) createSite(w http.ResponseWriter, r *http.Request) {
	site := &pb.Site{}
	err := json.NewDecoder(r.Body).Decode(site)
	if err != nil {
		return
	}
	if site.Url == "" {
		http.Error(w, "url is not defined", http.StatusBadRequest)
		return
	}
	_, err = s.hand.CreateAccount(&models.Account{
		URL: site.Url,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to create account: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *server) makeHandlers() {
	s.router.HandleFunc("/v1/users", s.createAccount)
}

func (s *server) startServer() {
	log.Infof("server is started at %s", s.address)
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
		hand:    h,
		router:  r,
		address: c.Address,
	}
	s.makeHandlers()
	s.startServer()
}
