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
		writeErrorResponse(w, fmt.Sprintf("unable to get stats: %v", err))
		return
	}

	site, ok := r.URL.Query()["site"]
	if !ok {
		writeErrorResponse(w, "site id is not defined")
		return
	}

	parsedSite, err := strconv.ParseInt(site[0], 10, 64)
	if err != nil {
		writeErrorResponse(w, err.Error())
		return
	}

	resp, err := s.hand.GetStats(&pb.GetStatsRequest{
		SiteID: parsedSite,
	})
	if err != nil {
		writeErrorResponse(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(resp)
	if err != nil {
		writeErrorResponse(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(data))
}

func (s *server) aggregateStats(w http.ResponseWriter, r *http.Request) {
	siteID, ok := r.URL.Query()["site"]
	if !ok {
		writeErrorResponse(w, "site id is not defined")
		return
	}

	parsedSite, err := strconv.ParseInt(siteID[0], 10, 64)
	if err != nil {
		writeErrorResponse(w, err.Error())
		return
	}

	resp, err := s.hand.AggregateStats(&pb.CountStatRequest{
		SiteID: parsedSite,
	})
	if err != nil {
		writeErrorResponse(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(resp)
	if err != nil {
		writeErrorResponse(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(data))
}
