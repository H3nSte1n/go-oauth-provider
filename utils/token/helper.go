package token

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"os"
)

func verify(tokenString string) (*jwt.Token, error) {
  token, err := jwt.Parse(tokenString, jwtParseCallback)
  if err != nil {
     return nil, err
  }

	if token.Valid {
		return token, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, fmt.Errorf("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, fmt.Errorf("Token is either expired or not active yet")
		} else {
			return nil, fmt.Errorf("Couldn't handle this error")
		}
	} else {
		return nil, fmt.Errorf("Couldn't handle this token")
	}
}

func jwtParseCallback(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		 return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(os.Getenv("SIGNATURE_SECRET")), nil
}