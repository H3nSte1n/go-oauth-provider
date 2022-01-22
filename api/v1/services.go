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

type ServiceList struct {
	Scopes []models.Service `json:"services" binding:"dive"`
}

func CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	currentTime := time.Now()
	service.CreatedAt = &currentTime
	id, err := db.CreateService(&service)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	
	c.JSON(200, gin.H{"id": id})
}

func UpdateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	currentTime := time.Now()
	service.UpdatedAt = &currentTime
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	id, err := db.UpdateService(docID, &service)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetService(c *gin.Context) {
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	service, err := db.GetService(docID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{"service": service})
}

func GetServices(c *gin.Context) {
	services, err := db.GetServices()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"services": services,
	})
}