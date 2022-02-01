package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccessGroup struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" binding:"required,min=3,max=255"`
	Resources []string           `json:"resources"`
	CreatedAt *time.Time         `json:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at"`
}
