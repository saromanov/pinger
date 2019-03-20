package storage

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
)

// GetSitesRequest provides argument with params
// for find sites
type GetSitesRequest struct {
	Limit  uint
	Offset uint
	UserID string
}

// GetSites returns list of sites
func (s *Storage) GetSites(req GetSitesRequest) ([]*models.Site, error) {
	var sites []*models.Site
	err := s.db.Where(nil).Find(&sites).Error
	if err != nil {
		return nil, errors.Wrap(err, "storage: unable to find sites")
	}
	return sites, nil
}
