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

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/saromanov/pinger/config"
	"github.com/saromanov/pinger/internal/handler"
	"github.com/saromanov/pinger/internal/log"
	"github.com/saromanov/pinger/internal/models"
	pb "github.com/saromanov/pinger/proto"
)

var tokenAuth *jwtauth.JWTAuth

type server struct {
	hand    *handler.Handler
	router  *chi.Mux
	address string
}

// createAccount makes a new account
func (s *server) createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	account := &pb.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		return
	}
	token, id, err := s.hand.CreateAccount(&models.Account{
		Email:    account.Email,
		Password: account.Password,
		Name:     account.Name,
	})
	if err != nil {
		writeResponse(w, ErrorResponse{
			Message: fmt.Sprintf("unable to create account: %v", err),
			Status:  "error",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, AccountResponse{
		ID:          fmt.Sprintf("%d", id),
		Token:       token,
		CreatedTime: time.Now().UTC(),
	})
}

// getAccount returns account by id
func (s *server) getAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, err := s.getUserFromContextToken(r.Context())
	if err != nil {
		writeResponse(w, ErrorResponse{
			Message: fmt.Sprintf("unable to validate token: %v", err),
			Status:  "error",
		})
		return
	}

	acc, err := s.hand.GetAccount(userID, "")
	if err != nil {
		writeResponse(w, ErrorResponse{
			Message: fmt.Sprintf("unable to get account: %v", err),
			Status:  "error",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponse(w, AccountResponse{
		ID:          fmt.Sprintf("%d", acc.ID),
		CreatedTime: time.Now().UTC(),
		Name:        acc.Name,
		Email:       acc.Email,
	})
}


func (s *server) makeHandlers() {
	s.router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/v1/users/{id}", s.getAccount)
		r.Post("/v1/sites", s.createSite)
		r.Get("/v1/stats", s.getStats)
	})

	s.router.Group(func(r chi.Router) {
		r.Post("/v1/users", s.createAccount)
	})
}

// getUserFromContextToken returns user id based on jwt token
// from context if token is not provides or invalid, it returns error
func (s *server) getUserFromContextToken(con context.Context) (int, error) {
	_, claims, err := jwtauth.FromContext(con)
	if err != nil {
		return 0, fmt.Errorf("unable to get token: %v", err)
	}
	user, ok := claims["UserID"]
	if !ok {
		return 0, fmt.Errorf("unable to get user id from claims")
	}

	return int(user.(float64)), nil

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
	tokenAuth = jwtauth.New("HS256", []byte(c.Token), nil)
	r := chi.NewRouter()
	s := &server{
		hand:    h,
		router:  r,
		address: c.Address,
	}
	s.makeHandlers()
	s.startServer()
}
