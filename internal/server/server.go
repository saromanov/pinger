// Package server defines rest api endpoints for pinger
package server

import (
	"context"
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
)

var tokenAuth *jwtauth.JWTAuth

type server struct {
	hand    *handler.Handler
	router  *chi.Mux
	address string
}

func (s *server) makeHandlers() {
	s.router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/v1/users/{id:[0-9]+}", s.getAccount)
		r.Get("/v1/me", s.me)
		r.Post("/v1/sites", s.createSite)
		r.Get("/v1/stats", s.getStats)
		r.Delete("/v1/stats/{id}", s.deleteSite)
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

// writeErrorResponse writes error response
func writeErrorResponse(w http.ResponseWriter, msg string) {
	writeResponse(w, ErrorResponse{
		Message: msg,
		Status:  "error",
	})
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
