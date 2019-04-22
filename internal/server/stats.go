package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/saromanov/pinger/internal/models"
	pb "github.com/saromanov/pinger/proto"
)

// getStats returns site statictics
func (s *server) getStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID, err := s.getUserFromContextToken(r.Context())
	if err != nil {
		writeResponse(w, ErrorResponse{
			Message: fmt.Sprintf("unable to get stats: %v", err),
			Status:  "error",
		})
		return
	}
	
	w.WriteHeader(http.StatusOK)
}
