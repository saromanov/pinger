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
	if account.Email == "" && account.Password == "" {
		http.Error(w, "email or password is not defined", http.StatusBadRequest)
		return
	}
	token, id, err := s.hand.CreateAccount(&models.Account{
		Email:    account.Email,
		Password: account.Password,
		Name:     account.Name,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to create account: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, AccountResponse{
		ID:          fmt.Sprintf("%d", id),
		Token:       token,
		CreatedTime: time.Now().UTC(),
	})
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
	_, err = s.hand.CreateSite(&models.Site{
		URL: site.Url,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to create account: %v", err), http.StatusBadRequest)
		return
	}

	resp := &AccountResponse{
		CreatedTime: time.Now().UTC(),
		ID:          "",
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal data: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(data))
	w.WriteHeader(http.StatusCreated)
}

func (s *server) makeHandlers() {
	s.router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/v1/users/{id}", func(w http.ResponseWriter, r *http.Request) {
			tok, claims, err := jwtauth.FromContext(r.Context())
			fmt.Println(tok, claims, err)
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})

	s.router.Group(func(r chi.Router) {
		r.Post("/v1/users", s.createAccount)
	})
}

// getUserFromContextToken returns user id based on jwt token
// from context if token is not provides or invalid, it returns error
func (s *server) getUserFromContextToken(con context.Context) (string, error) {
	_, claims, err := jwtauth.FromContext(con)
	if err != nil {
		return "", fmt.Errorf("unable to get token: %v", err)
	}

	user, ok := claims["UserId"]
	if !ok {
		return "", fmt.Errorf("unable to get user id from claims")
	}

	return user.(string), nil

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
	tokenAuth = jwtauth.New("HS256", []byte("testtoken"), nil)
	r := chi.NewRouter()
	s := &server{
		hand:    h,
		router:  r,
		address: c.Address,
	}
	s.makeHandlers()
	s.startServer()
}
