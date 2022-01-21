package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Credential struct {
	ID 							primitive.ObjectID  	`json:"id" bson:"_id,omitempty"`
	ClientSecret 		string 								`json:"client_secret"`
	ClientId 				string 								`json:"client_id"`
	CreatedAt 			time.Time 						`json:"created_at"`
	UpdatedAt 			time.Time 						`json:"updated_at"`
	ScopeIDs  			[]primitive.ObjectID  `json:"scopes"`
}