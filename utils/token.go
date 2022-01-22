package utils

import (
	"github.com/dgrijalva/jwt-go"
	// "os"
	// "log"
	"oauth_provider/models"
	"time"
)

type accessTokenClaims struct {
	models.Scope
	jwt.StandardClaims
}

func CreateAccessToken(ownClaims models.Scope) *jwt.Token {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := accessTokenClaims{
		ownClaims,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken
}

func SignToken(token *jwt.Token) (*jwt.Token, error) {
	// signature := os.Getenv("SIGNATURE_SECRET")
	// token, err := token.SignedString([]byte(signature))

	// if err != nil {
	// 	log.Print(err)
	// 	return nil, err
	// }
	return token, nil
}