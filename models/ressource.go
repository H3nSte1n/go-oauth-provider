package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Avaiable Ressources
type Ressource struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" binding:"required,min=5,max=255"`
	Url       string             `json:"url" binding:"required,url`
	CreatedAt *time.Time         `json:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at"`
}
