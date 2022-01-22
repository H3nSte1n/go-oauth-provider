package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"oauth_provider/models"
	"oauth_provider/db"
	"oauth_provider/utils"
	"time"
)

type ScopeList struct {
	Scopes []models.Scope `json:"scopes" binding:"dive"`
}

func CreateCredentials(c *gin.Context) {
	clientSecret := uuid.New().String()
	clientId := uuid.New().String()

	var scopes ScopeList
	if err := c.ShouldBindJSON(&scopes); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	convertedScopes := utils.ConvertToInterfaceArray[models.Scope](scopes.Scopes)
	scopeIds, _ := db.CreateScopes(convertedScopes)
	currentTime := time.Now()
	credential := models.Credential{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		ScopeIDs:     scopeIds,
		CreatedAt:    &currentTime,
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