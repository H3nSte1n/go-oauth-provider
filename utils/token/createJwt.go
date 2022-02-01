package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtTokenClaims[T any] struct {
	OwnClaims T `json:"OwnClaims"`
	jwt.StandardClaims
}

func CreateJwt[T any](ownClaims T) *jwt.Token {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := JwtTokenClaims[T]{
		ownClaims,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "oauth_provider",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken
}
