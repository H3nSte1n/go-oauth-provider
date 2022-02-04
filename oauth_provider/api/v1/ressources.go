package v1

import (
	"log"
	"net/http"
	"oauth_provider/db"
	"oauth_provider/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

	c.JSON(http.StatusOK, gin.H{"id": id})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"ressources": ressources,
	})
}
