package storage

import (
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
	pb "github.com/saromanov/pinger/proto"
)

var errNoSite = errors.New("site id is not defined")

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
func GetStats(req *pb.GetStatsRequest)([]*models.Ping, error) {
	if req.SiteID == 0 {
		return nil, errNoSite
	}

	r := &models.Ping{
		SiteID: req.SiteID,
		UserID: req.UserID,
	}
	var pings []*models.Ping
	if err := s.db.Where(r).Find(&pings).Error; err != nil {
		return nil, fmt.Errorf("unable to find stats: %v", err)
	}
	return pings, nil
}
