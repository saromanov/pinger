package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	pb "github.com/saromanov/pinger/proto"
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

	site, ok := r.URL.Query()["site"]
	if !ok {
		writeResponse(w, ErrorResponse{
			Message: "site id is not defined",
			Status:  "error",
		})
		return
	}

	parsedSite, err := strconv.ParseInt(site[0], 10, 64)
	if err != nil {
		writeResponse(w, ErrorResponse{
			Message: err.Error(),
			Status:  "error",
		})
		return
	}

	resp, err := s.hand.GetStats(&pb.GetStatsRequest{
		SiteID: parsedSite,
	})
	if err != nil {
		writeResponse(w, ErrorResponse{
			Message: err.Error(),
			Status:  "error",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal data: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(data))
}
