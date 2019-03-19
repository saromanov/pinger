package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// Account represents user's account
type Account struct {
	gorm.Model
	Email     string
	Password  string
	CreatedAt time.Time
	Token string
	Name string
}

// Token represents jwt token
type Token struct {
	UserID uint
	jwt.StandardClaims
}
