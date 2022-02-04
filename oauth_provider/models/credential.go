package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credential struct {
	ID           primitive.ObjectID   `json:"_id,omitempty"`
	ClientSecret *string              `json:"client_secret"`
	ClientID     *string              `json:"client_id"`
	CreatedAt    *time.Time           `json:"created_at"`
	UpdatedAt    *time.Time           `json:"updated_at"`
	ScopeIDs     []primitive.ObjectID `json:"scopes"`
}
