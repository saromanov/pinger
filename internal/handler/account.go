package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateAccount provides creating of the new user
// Its generate a new bassword with bcrypt library
// and then, add to the storage
func (h *Handler) CreateAccount(u *models.Account) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "unable to hash password")
	}
	u.Password = string(pass)

	if err := h.Storage.InsertAccount(u); err != nil {
		return errors.Wrap(err, "unable to create account")
	}
	return nil
}

// creating of the jwt token
func createJWTToken(u *models.Account) *models.Account {
	tk := &models.Token{UserID: u.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("testtoken"))
	u.Token = tokenString
	return u
}
