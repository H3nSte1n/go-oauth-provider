package token

import "github.com/dgrijalva/jwt-go"

func Valid(tokenString string) error {
  token, err := verify(tokenString)

  if err != nil {
     return err
  }
  if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
     return err
  }
  return nil
}