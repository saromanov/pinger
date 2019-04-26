package handler

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
	"github.com/saromanov/pinger/internal/storage"
	pb "github.com/saromanov/pinger/proto"
)

var errNoUrl = errors.New("url is not defined")

// CreateSite provides creating of the new site
// for checking of availability. It should be attached to user
func (h *Handler) CreateSite(u *models.Site) (string, error) {
	if u.URL == "" {
		return "", errNoUrl
	}

	_, err := url.ParseRequestURI(u.URL)
	if err != nil {
		return "", fmt.Errorf("unable to parse url: %v", err)
	}

	if err = h.Storage.InsertSite(u); err != nil {
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

// GetSite returns site by id
func (h *Handler) GetSite(r *pb.GetSiteRequest) (*pb.Site, error) {
	site, err := h.Storage.GetSite(r.Id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get site")
	}

	return convertSiteToProto(site), nil
}

func (h *Handler) DeleteSite(s int64) error {
	err := h.Storage.DeleteSite(s)
	if err != nil {
		return errors.Wrap(err, "unable to get site")
	}

	return nil
}

func convertSitesToProto(sites []*models.Site) []*pb.Site {
	result := make([]*pb.Site, len(sites))
	for i, s := range sites {
		result[i] = convertSiteToProto(s)
	}
	return result
}

func convertSiteToProto(s *models.Site) *pb.Site {
	return &pb.Site{
		Url: s.URL,
		Id:  int64(s.ID),
	}
}
