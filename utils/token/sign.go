package token

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"log"
)

func Sign(token *jwt.Token) (*string, error) {
	signature := os.Getenv("SIGNATURE_SECRET")
	signedToken, err := token.SignedString([]byte(signature))

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &signedToken, nil
}