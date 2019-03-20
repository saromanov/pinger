// Package handler contains internal logic of the app
// Its getting models and validate it and store
package handler

import (
	"errors"
	"github.com/saromanov/pinger/internal/storage"
)

var errNoStorage = errors.New("storage is not defined")

// Handler represents part for operation with db
type Handler struct {
	Storage *storage.Storage
}

// New provides initialization of the handler
func New(s *storage.Storage) (*Handler, error) {
	if s == nil {
		return nil, errNoStorage
	}

	return &Handler{
		Storage: s,
	}, nil
}
