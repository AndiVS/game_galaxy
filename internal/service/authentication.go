package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/AndiVS/game_galaxy/internal/model"
	"github.com/AndiVS/game_galaxy/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Authentication aunt
type Authentication interface {
	SignUp(c context.Context, account *model.Account) error
	SignIn(c context.Context, account *model.Account) (string, error)
}

// AuthenticationService aunt
type AuthenticationService struct {
	Rep      repository.Account
	Access   *JWTManager
	Refresh  *JWTManager
	HashSalt string
}

// NewServiceAuthentication create aunt
func NewServiceAuthentication(repositories interface{}, access, refresh *JWTManager, hashSalt string) Authentication {
	return &AuthenticationService{Rep: repositories.(*repository.Postgres), Access: access, Refresh: refresh, HashSalt: hashSalt}
}

// SignUp record about cat
func (s *AuthenticationService) SignUp(c context.Context, account *model.Account) error {
	account.Password = PasswordGenerator(account.Password, s.HashSalt)

	_, err := s.Rep.SelectAccount(c, account.Login)
	if err != nil {
		return s.Rep.InsertAccount(c, account)
	}

	return echo.NewHTTPError(http.StatusInternalServerError, "UNABLE TO INSERT ")
}

// SignIn generate token
func (s *AuthenticationService) SignIn(c context.Context, account *model.Account) (accessToken string, err error) {
	account.Password = PasswordGenerator(account.Password, s.HashSalt)

	accountFromBase, err := s.Rep.SelectAccount(c, account.Login)
	if err != nil {
		return "", err
	}
	if !PasswordCheck(accountFromBase.Password, account.Password) {
		return "", err
	}

	return GenerateTokens(account, s.Access)
}

// PasswordCheck che
func PasswordCheck(pass1, pass2 string) bool {
	return pass1 == pass2
}

// PasswordGenerator generate password from hash and pass string
func PasswordGenerator(password, hashSalt string) string {
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(hashSalt))

	return fmt.Sprintf("%x", pwd.Sum(nil))
}
