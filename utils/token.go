package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"log"
	"time"
)

type accessTokenClaims[T any] struct {
	OwnClaims T
	jwt.StandardClaims
}

func CreateAccessToken[T any](ownClaims T) *jwt.Token {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := accessTokenClaims[T]{
		ownClaims,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "oauth_provider",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken
}

func SignToken(token *jwt.Token) (*string, error) {
	signature := os.Getenv("SIGNATURE_SECRET")
	signedToken, err := token.SignedString([]byte(signature))

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &signedToken, nil
}