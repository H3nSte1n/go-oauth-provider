package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Scope struct {
	ID          primitive.ObjectID  `json:"_id,omitempty"`
	Name        string              `json:"name" binding:"required,min=5,max=255"`
	Read        *bool               `json:"read" binding:"required"`
	Write       *bool               `json:"write" binding:"required"`
	Create      *bool               `json:"create" binding:"required"`
	Delete      *bool               `json:"delete" binding:"required"`
	RessourceID *primitive.ObjectID `json:"ressource_id" binding:"required"`
	CreatedAt   *time.Time          `json:"created_at"`
	UpdatedAt   *time.Time          `json:"updated_at"`
}
