package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Credential struct {
	ID 							primitive.ObjectID  	`json:"id" bson:"_id,omitempty"`
	ClientSecret 		string 								`json:"client_secret"`
	ClientId 				string 								`json:"client_id"`
}