package handler

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
	pb "github.com/saromanov/pinger/proto"
)

// GetSites returns list of the sites based on request
func (h *Handler) GetSites(u *models.Account) ([]*pb.Site, error) {
	sites, err := h.Storage.GetSites()
	if err != nil {
		return nil, errors.Wrap(err, "unable to get sites")
	}

	return convertSitesToProto(sites), nil
}

func convertSitesToProto(sites []*models.Site) []*pb.Site {
	result := make([]*pb.Site, len(sites))
	for i, s := range sites {
		result[i] = &pb.Site{
			Url: s.URL,
			Id:  int64(s.ID),
		}
	}
	return result
}
