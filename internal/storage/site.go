package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
)

var (
	errNoSite = errors.New("site is not defined")
)

// GetSitesRequest provides argument with params
// for find sites
type GetSitesRequest struct {
	Limit  int
	Offset int
	UserID string
}

// InsertSite provides inserting of site
func (s *Storage) InsertSite(m interface{}) (uint, error) {
	resp := &models.Site{}
	err := s.db.Create(m).Scan(resp).Error
	if err != nil {
		return 0, errors.Wrap(err, "storage: unable to insert site")
	}
	return resp.ID, nil
}

// GetSites returns list of sites
func (s *Storage) GetSites(req GetSitesRequest) ([]*models.Site, error) {
	if req.UserID == "" {
		return s.getAllSites(req)
	}
	return s.getUserSites(req)
}

// GetSite return site by id
func (s *Storage) GetSite(id int64) (*models.Site, error) {
	if id == 0 {
		return nil, errNoSite
	}

	var site *models.Site
	if err := s.db.Where("id = ?", id).First(&site).Error; err != nil {
		return nil, fmt.Errorf("unable to find site: %v", err)
	}
	return site, nil
}

// DeleteSite provides removing of the site
func (s *Storage) DeleteSite(id int64) error {
	if id == 0 {
		return errNoSite
	}

	if err := s.db.Delete(&models.Site{
		Model: gorm.Model{
			ID: uint(id),
		},
	}).Error; err != nil {
		return fmt.Errorf("unable to delete site: %v", err)
	}

	return nil
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
