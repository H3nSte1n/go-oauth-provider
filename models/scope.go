package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Scope struct {
	ID 						primitive.ObjectID  	`json:"id" bson:"_id,omitempty"`
	Name 					string 								`json:"name" binding:"required,min=5,max=255"`
	Read 					*bool 								`json:"read" binding:"required"`
	Write 				*bool 								`json:"write" binding:"required"`
	Create 				*bool 								`json:"create" binding:"required"`
	Delete 				*bool 								`json:"delete" binding:"required"`
	ServiceID     *primitive.ObjectID   `json:"service_id" binding:"required"`
	CreatedAt 		*time.Time 						`json:"created_at"`
	UpdatedAt 		*time.Time 						`json:"updated_at"`
}