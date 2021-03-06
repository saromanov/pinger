package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/saromanov/pinger/internal/models"
	pb "github.com/saromanov/pinger/proto"
)

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
	id := chi.URLParam(r, "id")
	if id == "" {
		writeErrorResponse(w, "id is not defined")
		return
	}
	_, err := s.getUserFromContextToken(r.Context())
	if err != nil {
		writeResponse(w, ErrorResponse{
			Message: fmt.Sprintf("unable to validate token: %v", err),
			Status:  "error",
		})
		return
	}

	parsedInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		writeErrorResponse(w, err.Error())
		return
	}
	acc, err := s.hand.GetAccount(int(parsedInt), "")
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

func (s *server) me(w http.ResponseWriter, r *http.Request) {
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
