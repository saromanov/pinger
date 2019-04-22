package server

import (
	"fmt"
	"net/http"
)

// getStats returns site statictics
func (s *server) getStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := s.getUserFromContextToken(r.Context())
	if err != nil {
		writeResponse(w, ErrorResponse{
			Message: fmt.Sprintf("unable to get stats: %v", err),
			Status:  "error",
		})
		return
	}

	site, ok := r.URL.Query()["site_id"]
	if !ok {
		writeResponse(w, ErrorResponse{
			Message: "site id is not defined",
			Status:  "error",
		})
		return
	}

	fmt.Println(site)
	w.WriteHeader(http.StatusOK)
}
