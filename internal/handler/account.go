package handler

import (
	"fmt"

	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/saromanov/pinger/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	errNoPassword    = errors.New("password is not defined")
	errNoEmail       = errors.New("email is not defined")
	errAccountExist  = errors.New("account with such email already exist")
	errAccountNoName = errors.New("account not have a name")
)

// CreateAccount provides creating of the new user
// Its generate a new bassword with bcrypt library
// and then, add to the storage
func (h *Handler) CreateAccount(u *models.Account) (string, uint, error) {
	if err := validateCreds(u); err != nil {
		return "", 0, err
	}
	if _, err := h.Storage.GetAccount(0, u.Email); err == nil {
		return "", 0, errAccountExist
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", 0, errors.Wrap(err, "unable to hash password")
	}
	u.Password = string(pass)

	id, err := h.Storage.InsertAccount(u)
	if err != nil {
		return "", 0, errors.Wrap(err, "unable to create account")
	}
	return createJWTToken(u), id, nil
}

// GetAccount returns account by id
func (h *Handler) GetAccount(id int, email string) (*models.Account, error) {
	acc, err := h.Storage.GetAccount(id, email)
	if err != nil {
		return nil, fmt.Errorf("unable to get account: %v", err)
	}
	return acc, nil
}

// Login provides auth for the user
// It contains validation of email and generating of password hash
func (h *Handler) Login(email, password string) (*models.Account, error) {
	acc, err := h.Storage.GetAccount(0, email)
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

// validateCreds provides validation of the input data
// for account
func validateCreds(u *models.Account) error {
	if err := validateEmail(u.Email); err != nil {
		return err
	}
	if u.Password == "" {
		return errNoPassword
	}
	if u.Name == "" {
		return errAccountNoName
	}
	return nil
}

// validateEmail provides validation of email format
func validateEmail(email string) error {
	if email == "" {
		return errNoEmail
	}
	if err := checkmail.ValidateFormat(email); err != nil {
		return errors.Wrap(err, "unable to validate email format")
	}

	return nil
}

// creating of the jwt token
func createJWTToken(u *models.Account) string {
	tk := &models.Token{UserID: u.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("testtoken"))
	return tokenString
}
