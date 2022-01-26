package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AccessGroup struct {
	ID        primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	Name      string 								`json:"name" binding:"required,min=3,max=255"`
	Resources []primitive.ObjectID  `json:"resources"`
	CreatedAt *time.Time 						`json:"created_at"`
	UpdatedAt *time.Time 						`json:"updated_at"`
}