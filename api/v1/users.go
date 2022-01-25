package v1

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"log"
	"time"

	"oauth_provider/models"
	"oauth_provider/db"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	bytes, bycrypt_err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	currentTime := time.Now()
	user.CreatedAt = &currentTime
	id, err := db.CreateUser(&user)

	if err != nil || bycrypt_err != nil {
		a := err 
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error: " + a.Error() + " Bycrypt: " + bycrypt_err.Error()})
		return
	}
	
	c.JSON(200, gin.H{"id": id})
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
	user_id, err := db.DeleteUser(docID)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": user_id,
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
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}