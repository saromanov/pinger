package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Ping represents model for storage data
// of the site availability
type Ping struct {
	gorm.Model
	ResponseTime time.Duration
	ResponseCode int
	Available    bool
	SiteID       int64
	UserID       string
}
