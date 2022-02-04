package middleware

import (
	"net/http"
	"oauth_provider/utils/token"

	"github.com/gin-gonic/gin"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := token.Extract(c.Request)
		err := token.Valid(c.Request.URL.Path[1:], tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()

			return
		}

		c.Next()
	}
}
