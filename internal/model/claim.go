package model

import "github.com/golang-jwt/jwt"

// Claims for JWT
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
