package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Scope struct {
	ID 						primitive.ObjectID  	`json:"id" bson:"_id,omitempty"`
	Name 					string 								`json:"name"`
	Read 					bool 									`json:"read"`
	Write 				bool 									`json:"write"`
	Create 				bool 									`json:"create"`
	Delete 				bool 									`json:"delete"`
	CreatedAt 		time.Time 						`json:"created_at"`
	UpdatedAt 		time.Time 						`json:"updated_at"`
}