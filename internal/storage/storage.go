// Package storage contains handling with db(Postgesql)
package storage

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saromanov/pinger/config"
	"github.com/saromanov/pinger/internal/models"
)

var errNoConfig = errors.New("config is not defined")

// Storage implements db handling with Postgesql
type Storage struct {
	db *gorm.DB
}

// New provides init for postgesql storage
func New(s *config.Config) (*Storage, error) {
	if s == nil {
		return nil, errNoConfig
	}
	args := "dbname=pinger"
	if s.Name != "" && s.Password != "" && s.User != "" {
		args += fmt.Sprintf(" user=%s dbname=%s password=%s", s.User, s.Name, s.Password)
	}
	db, err := gorm.Open("postgres", args)
	if err != nil {
		return nil, fmt.Errorf("unable to open db: %v", err)
	}
	db.AutoMigrate(&models.Site{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Ping{})
	return &Storage{
		db: db,
	}, nil
}

// Close provides closing of db
func (s *Storage) Close() error {
	return s.db.Close()
}
