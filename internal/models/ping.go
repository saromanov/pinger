package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Ping represents model for storage data
// of the site availability
type Ping struct {
	gorm.Model
	ID           int64
	ResponseTime time.Time
	ResponseCode int
	Available    bool
}