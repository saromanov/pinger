package models

import (
	"github.com/jinzhu/gorm"
)

// PingData represents model for storage data
// of the site availability
type PingData struct {
	gorm.Model
	ResponseTime int64
	ResponseCode int
	Available    bool
	SiteID       int64
	UserID       string
	ErrorMessage string
}
