package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/saromanov/pinger/internal/models"
	pb "github.com/saromanov/pinger/proto"
)

// createSite makes a new site for user
func (s *server) createSite(w http.ResponseWriter, r *http.Request) {
	userID, err := s.getUserFromContextToken(r.Context())
	if err != nil {
		writeResponse(w, ErrorResponse{
			Message: fmt.Sprintf("unable to create site: %v", err),
			Status:  "error",
		})
		return
	}
	site := &pb.Site{}
	err = json.NewDecoder(r.Body).Decode(site)
	if err != nil {
		return
	}
	if site.Url == "" {
		http.Error(w, "url is not defined", http.StatusBadRequest)
		return
	}
	id, err := s.hand.CreateSite(&models.Site{
		URL:    site.Url,
		UserID: fmt.Sprintf("%d", userID),
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to create site: %v", err), http.StatusBadRequest)
		return
	}

	resp := &AccountResponse{
		CreatedTime: time.Now().UTC(),
		ID:          id,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal data: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(data))
	w.WriteHeader(http.StatusCreated)
}
