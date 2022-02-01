package middleware

import (
	"github.com/gin-gonic/gin"

	"oauth_provider/utils/token"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := token.Extract(c.Request)
		err := token.Valid(c.Request.URL.Path[1:], tokenString)

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
