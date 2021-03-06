package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oauth_provider/models"
)

var USER_MODEL = "users"

func CreateUser(user *models.User) (primitive.ObjectID, error) {
	user.ID = primitive.NewObjectID()
	return Create(user, USER_MODEL)
}

func UpdateUser(id primitive.ObjectID, user *models.User) (primitive.ObjectID, error) {
	return Update(USER_MODEL, id, user)
}

func DeleteUser(id primitive.ObjectID) (*primitive.ObjectID, error) {
	return Delete[*models.User](USER_MODEL, id)
}

func GetUsers() ([]models.User, error) {
	return GetList[models.User](USER_MODEL)
}

func GetUser(id primitive.ObjectID) (models.User, error) {
	return Get[models.User](USER_MODEL, id)
}

func UserFindByUsername(username *string) (models.User, error) {
	attributes := map[string]interface{}{
		"username": username,
	}

	return FindByAttributes[models.User](USER_MODEL, attributes)
}
