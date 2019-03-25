package handler

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
	"github.com/saromanov/pinger/internal/storage"
	pb "github.com/saromanov/pinger/proto"
)

// CreateSite provides creating of the new site
// for checking of availability. It should be attached to user
func (h *Handler) CreateSite(u *models.Site) (string, error) {
	if err := h.Storage.InsertSite(u); err != nil {
		return "", errors.Wrap(err, "unable to create site")
	}
	return "", nil
}

// GetSites returns list of the sites based on request
func (h *Handler) GetSites(r *pb.GetSitesRequest) ([]*pb.Site, error) {
	sites, err := h.Storage.GetSites(storage.GetSitesRequest{
		Limit:  int(r.Limit),
		Offset: int(r.Offset),
	})
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
