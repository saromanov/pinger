package handler

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
	pb "github.com/saromanov/pinger/proto"
)

// CreateStat provides creating of the new site
// for checking of availability. It should be attached to user
func (h *Handler) CreateStat(u *models.PingData) (string, error) {
	if _, err := h.Storage.InsertStat(u); err != nil {
		return "", errors.Wrap(err, "unable to create site")
	}
	return "", nil
}

// GetStats returns ping stat
func (h *Handler) GetStats(req *pb.GetStatsRequest) ([]*models.PingData, error) {
	resp, err := h.Storage.GetStats(req)
	if err != nil {
		return nil, errors.Wrap(err, "unable to search stats")
	}
	return resp, nil
}
