package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID 						primitive.ObjectID  	`json:"id" bson:"_id,omitempty"`
	FirstName 		string 								`json:"first_name" binding:"required,max=255"`
	LastName 			string 								`json:"last_name" binding:"required"`
	Email 				string 								`json:"email" binding:"required,email"`
	Username 			string 								`json:"username" binding:"required,min=5,max=255"`
	Password 			string 								`json:"password" binding:"required,min=5,max=255"`
	AccessGroups 	[]primitive.ObjectID 	`json:"access_groups" binding:"required"`
	CreatedAt 		*time.Time 						`json:"created_at"`
	UpdatedAt 		*time.Time 						`json:"updated_at"`
}