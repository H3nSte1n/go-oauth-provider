package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oauth_provider/models"
	"oauth_provider/db"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RessourceList struct {
	Scopes []models.Ressource `json:"ressources" binding:"dive"`
}

func CreateRessource(c *gin.Context) {
	var ressource models.Ressource
	if err := c.ShouldBindJSON(&ressource); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	currentTime := time.Now()
	ressource.CreatedAt = &currentTime
	id, err := db.CreateRessource(&ressource)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	
	c.JSON(200, gin.H{"id": id})
}

func UpdateRessource(c *gin.Context) {
	var ressource models.Ressource
	if err := c.ShouldBindJSON(&ressource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	currentTime := time.Now()
	ressource.UpdatedAt = &currentTime
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	id, err := db.UpdateRessource(docID, &ressource)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetRessource(c *gin.Context) {
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	ressource, err := db.GetRessource(docID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{"ressource": ressource})
}

func GetRessources(c *gin.Context) {
	ressources, err := db.GetRessources()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"ressources": ressources,
	})
}