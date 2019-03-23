package handler

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
	"github.com/saromanov/pinger/proto"
)

// GetSites returns list of the sites based on request
func (h *Handler) GetSites(u *models.Account) ([]*proto.Site, error) {
	sites, err := h.Storage.GetSites()
	if err != nil {
		return nil, errors.Wrap(err, "unable to get sites")
	}

	return convertSitesToProto(sites), nil
}

func convertSitesToProto(sites []*models.Site) []*proto.Site {
	result := make([]*proto.Site, len(sites))
	for i, s := range sites {
		result[i] = &proto.Site{
			Url: s.URL,
			Id: s.Id,
		}
	}
	return result
}
