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
	Name        string    `json:"name"`
	Email       string    `json:"email"`
}

// ErrorResponse defines response after errors
type ErrorResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"	`
}

// write response
func writeResponse(w http.ResponseWriter, obj interface{}) {
	res, _ := json.Marshal(obj)
	fmt.Fprintln(w, string(res))
}
