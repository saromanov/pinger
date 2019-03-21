package models

import "github.com/jinzhu/gorm"

// Site represents model representation for site
type Site struct {
	gorm.Model
	URL string
}
