package v1

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"oauth_provider/db"
	"oauth_provider/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	cost := 14
	bytes, bycryptErr := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	user.Password = string(bytes)
	currentTime := time.Now()
	user.CreatedAt = &currentTime
	id, err := db.CreateUser(&user)

	if err != nil || bycryptErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error: " + err.Error() + " Bycrypt: " + bycryptErr.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	currentTime := time.Now()
	user.UpdatedAt = &currentTime
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	id, err := db.UpdateUser(docID, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func DeleteUser(c *gin.Context) {
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	userID, err := db.DeleteUser(docID)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
	})
}

func GetUser(c *gin.Context) {
	docID, _ := primitive.ObjectIDFromHex(c.Param("id"))
	user, err := db.GetUser(docID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUsers(c *gin.Context) {
	users, err := db.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
