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
func (h *Handler) CreateAccount(u *models.Account) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "unable to hash password")
	}
	u.Password = string(pass)

	if err := h.Storage.InsertAccount(u); err != nil {
		return "", errors.Wrap(err, "unable to create account")
	}
	return createJWTToken(u), nil
}

// Login provides auth for the user
func (h *Handler) Login(email, password string) (*models.Account, error) {
	acc, err := h.Storage.GetAccount(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, errors.Wrap(err, "invalid login credentials")
		}

		return nil, errors.Wrap(err, "unknown error on compare password hash")
	}
	acc.Password = ""
	acc.Token = createJWTToken(acc)
	return acc, nil
}

// creating of the jwt token
func createJWTToken(u *models.Account) string {
	tk := &models.Token{UserID: u.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("testtoken"))
	return tokenString
}
