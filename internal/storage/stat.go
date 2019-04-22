package storage

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
)

// InsertStat provides inserting of the ping stat
func (s *Storage) InsertStat(m interface{}) (uint, error) {
	resp := &models.Ping{}
	err := s.db.Create(m).Scan(resp).Error
	if err != nil {
		return 0, errors.Wrap(err, "storage: unable to insert ping data")
	}
	return resp.ID, nil
}

// GetStats returns statictics of pings
func GetStats()([]*models.Ping, error) {
	var pings []*models.Ping
	if err := s.db.Where().Find(&pings).Error {
		return nil, fmt.Errorf("unable to find stats: %v", err)
	}
	return pings, nil
}
