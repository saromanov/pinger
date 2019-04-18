package handler

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
)

// CreateStat provides creating of the new site
// for checking of availability. It should be attached to user
func (h *Handler) CreateStat(u *models.Ping) (string, error) {
	if _, err := h.Storage.InsertStat(u); err != nil {
		return "", errors.Wrap(err, "unable to create site")
	}
	return "", nil
}
