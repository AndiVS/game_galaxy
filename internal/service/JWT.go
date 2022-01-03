package service

import (
	"fmt"
	"time"

	"github.com/AndiVS/game_galaxy/internal/model"
	"github.com/golang-jwt/jwt"
)

// JWTManager is a JSON web token manager
type JWTManager struct {
	SecretKey     []byte
	TokenDuration time.Duration
}

// NewJWTManager returns a new JWT manager
func NewJWTManager(secretKey []byte, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{SecretKey: secretKey, TokenDuration: tokenDuration}
}

// GenerateTokens func for token generation
func GenerateTokens(account *model.Account, access *JWTManager) (accessToken string, err error) {
	accessToken, err = GenerateToken(account, access)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// GenerateToken generate jwt token
func GenerateToken(account *model.Account, manager *JWTManager) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.TokenDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Username: account.Login,
	})

	tokenString, err := token.SignedString(manager.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verify verifies the access token string and return a user claim if the token is valid
func (manager *JWTManager) Verify(Token string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(
		Token,
		&model.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return manager.SecretKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
