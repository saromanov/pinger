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

// InsertSite provides inserting of site
func (s *Storage) InsertSite(m interface{}) error {
	err := s.db.Create(m).Error
	if err != nil {
		return errors.Wrap(err, "storage: unable to insert site")
	}
	return nil
}

// GetSites returns list of sites
func (s *Storage) GetSites(req GetSitesRequest) ([]*models.Site, error) {
	if req.UserID == "" {
		return s.getAllSites(req)
	}
	return s.getUserSites(req)
}

// getAllSites needs for inner logic for checking site availibility
// Its getting batch of sites at one time
func (s *Storage) getAllSites(req GetSitesRequest) ([]*models.Site, error) {
	var sites []*models.Site
	err := s.db.Limit(req.Limit).Offset(req.Offset).Find(&sites).Error
	if err != nil {
		return nil, errors.Wrap(err, "storage: unable to find sites")
	}
	return sites, nil
}

// getUserSites returns sites which user registered
func (s *Storage) getUserSites(req GetSitesRequest) ([]*models.Site, error) {
	var sites []*models.Site
	err := s.db.Where("user_id = ?", req.UserID).Find(&sites).Error
	if err != nil {
		return nil, errors.Wrap(err, "storage: unable to find sites")
	}
	return sites, nil
}
