package v1

import (
	"log"
	"net/http"
	"oauth_provider/db"
	"oauth_provider/models"
	"oauth_provider/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ScopeList struct {
	Scopes []models.Scope `json:"scopes" binding:"dive"`
}

func CreateCredentials(c *gin.Context) {
	clientSecret := uuid.New().String()
	clientID := uuid.New().String()

	var scopes ScopeList
	if err := c.ShouldBindJSON(&scopes); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})

		return
	}

	convertedScopes := utils.ConvertToInterfaceArray(scopes.Scopes)
	scopeIds, _ := db.CreateScopes(convertedScopes)
	currentTime := time.Now()
	credential := models.Credential{
		ClientID:     &clientID,
		ClientSecret: &clientSecret,
		ScopeIDs:     scopeIds,
		CreatedAt:    &currentTime,
	}

	_, credentialsErr := db.CreateCredential(&credential)

	if credentialsErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": credentialsErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"client_id":     clientID,
		"client_secret": clientSecret,
	})
}

func GetCredentials(c *gin.Context) {
	credentials, err := db.GetCredentials()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"credentials": credentials,
	})
}
