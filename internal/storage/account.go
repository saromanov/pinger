package storage

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
)

// InsertAccount provides inserting of account
func (s *Storage) InsertAccount(m interface{}) error {
	err := s.db.Create(m).Error
	if err != nil {
		return errors.Wrap(err, "storage: unable to insert user")
	}
	return nil
}

// GetAccount returns account by id
func (s *Storage) GetAccount(id string) (*models.Account, error) {
	var acc *models.Account
	err := s.db.Where("email = ?", id).First(&acc).Error
	if err != nil {
		return nil, errors.Wrap(err, "storage: unable to get user")
	}
	return acc, nil
}
