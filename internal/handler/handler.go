// Package handler contains internal logic of the app
// Its getting models and validate it and store
package handler

import "github.com/saromanov/pinger/internal/storage"

// Handler represents part for operation with db
type Handler struct {
	Storage storage.Storage
}
