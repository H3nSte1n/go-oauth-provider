package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"oauth_provider/models"
	"oauth_provider/db"
	"oauth_provider/utils"
)

type ScopeList struct {
	Scopes []models.Scope `binding:"required"`
}

func CreateCredentials(c *gin.Context) {
	clientSecret := uuid.New().String()
	clientId := uuid.New().String()

	var scopes ScopeList
	if err := c.ShouldBindJSON(&scopes); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	convertedScopes := utils.ConvertToInterfaceArray[models.Scope](scopes.Scopes)
	scopeIds, _ := db.CreateScopes(convertedScopes)
	credential := models.Credential{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		ScopeIDs:     scopeIds,
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