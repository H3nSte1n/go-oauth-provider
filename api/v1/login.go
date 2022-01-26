package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"

	"oauth_provider/utils/token"
	"oauth_provider/utils/verify"
	"oauth_provider/models"
)

type LoginType struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var login LoginType
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	user, err := verify.User(&login.Username, &login.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err})
		return
	}

	tokenString := token.CreateJwt[models.User](*user)
	signedToken, err := token.Sign(tokenString)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}