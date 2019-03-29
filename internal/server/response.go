package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AccountResponse defines response
// after creating of account
type AccountResponse struct {
	ID          string    `json:"id"`
	Token       string    `json:"token"`
	CreatedTime time.Time `json:"time"`
}

// write response
func writeResponse(w http.ResponseWriter, obj interface{}) {
	res, _ := json.Marshal(obj)
	fmt.Fprintln(w, string(res))
}
