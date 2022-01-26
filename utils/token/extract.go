package token

import (
	"net/http"
	"strings"
)

func Extract(r *http.Request) string {
  token := r.Header.Get("Authorization")

	authArray := strings.Split(token, " ")
  if len(authArray) > 1 {
     return authArray[1]
  }
  return ""
}