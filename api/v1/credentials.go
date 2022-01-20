package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"oauth_provider/models"
	"oauth_provider/db"
	"fmt"
)

func CreateCredential(c *gin.Context) {
	clientSecret := uuid.New().String()
	clientId := uuid.New().String()

	credential := models.Credential{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
	
	db.CreateCredential(&credential)

	c.JSON(200, gin.H{
		"client_id": clientId,
		"client_secret": clientSecret,
	})
}


func GetCredentials(c *gin.Context) {
	credentials, err := db.GetCredentials()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"credentials": credentials,
	})
}