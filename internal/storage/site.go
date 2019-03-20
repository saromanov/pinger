package storage

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
)

// GetSitesRequest provides argument with params
// for find sites
type GetSitesRequest struct {
	Limit  int
	Offset int
	UserID string
}

// GetSites returns list of sites
func (s *Storage) GetSites(req GetSitesRequest) ([]*models.Site, error) {
	if req.UserID == "" {
		return s.getAllSites(req)
	}
	var sites []*models.Site
	err := s.db.Where("user_id = ?", req.UserID).Find(&sites).Error
	if err != nil {
		return nil, errors.Wrap(err, "storage: unable to find sites")
	}
	return sites, nil
}

func (s *Storage) getAllSites(req GetSitesRequest) ([]*models.Site, error) {
	var sites []*models.Site
	err := s.db.Limit(req.Limit).Offset(req.Offset).Find(&sites).Error
	if err != nil {
		return nil, errors.Wrap(err, "storage: unable to find sites")
	}
	return sites, nil
}
