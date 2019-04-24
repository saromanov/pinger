package models

import (
	"github.com/jinzhu/gorm"
)

// Ping represents model for storage data
// of the site availability
type Ping struct {
	gorm.Model
	ResponseTime int64
	ResponseCode int
	Available    bool
	SiteID       int64
	UserID       string
}
