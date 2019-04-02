package storage

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
)

// InsertAccount provides inserting of account
func (s *Storage) InsertAccount(m interface{}) (uint, error) {
	resp := &models.Account{}
	err := s.db.Create(m).Scan(resp).Error
	if err != nil {
		return 0, errors.Wrap(err, "storage: unable to insert user")
	}
	return resp.ID, nil
}

// GetAccount returns account by id
func (s *Storage) GetAccount(id int, email string) (*models.Account, error) {
	req := &models.Account{}
	if id != 0 {
		req.ID = uint(id)
	}
	if email != "" {
		req.Email = email
	}
	acc := &models.Account{}
	err := s.db.Where(req).First(&acc).Error
	if err != nil {
		return nil, errors.Wrap(err, "storage: unable to get user")
	}
	return acc, nil
}
