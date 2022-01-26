package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtTokenClaims[T any] struct {
	OwnClaims T
	jwt.StandardClaims
}

func CreateJwt[T any](ownClaims T) *jwt.Token {
	expirationTime := time.Now().Add(5 * time.Minute)
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