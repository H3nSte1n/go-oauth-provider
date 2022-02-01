package v1

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oauth_provider/db"
	"oauth_provider/models"
)

func CreateAccessGroup(c *gin.Context) {
	var accessGroup models.AccessGroup
	if err := c.ShouldBindJSON(&accessGroup); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})

		return
	}

	currentTime := time.Now()
	accessGroup.CreatedAt = &currentTime

	_, accessGroupError := db.CreateAccessGroup(&accessGroup)

	if accessGroupError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": accessGroupError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_group": accessGroup,
	})
}

func UpdateAccessGroup(c *gin.Context) {
	var accessGroup models.AccessGroup
	if err := c.ShouldBindJSON(&accessGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	currentTime := time.Now()
	accessGroup.UpdatedAt = &currentTime
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	id, err := db.UpdateAccessGroup(docID, &accessGroup)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func DeleteAccessGroup(c *gin.Context) {
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	accessGroupID, err := db.DeleteAccessGroup(docID)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"accessGroup_id": accessGroupID,
	})
}

func GetAccessGroup(c *gin.Context) {
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	accessGroup, err := db.GetAccessGroup(docID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{"accessGroup": accessGroup})
}

func GetAccessGroups(c *gin.Context) {
	accessGroups, err := db.GetAccessGroups()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"accessGroups": accessGroups,
	})
}
