package storage

import (
	"github.com/pkg/errors"
)

// InsertAccount provides inserting of account
func (s *Storage) InsertAccount(m interface{}) error {
	err := s.db.Create(m).Error
	if err != nil {
		return errors.Wrap(err, "storage: unable to insert user")
	}
	return nil
}
