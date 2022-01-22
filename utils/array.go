package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func Contains(items []primitive.ObjectID, item primitive.ObjectID) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}

	return false
}