package server

import "time"

// AccountResponse defines response
// after creating of account
type AccountResponse struct {
	ID          string    `json:"id"`
	CreatedTime time.Time `json:"time"`
}
