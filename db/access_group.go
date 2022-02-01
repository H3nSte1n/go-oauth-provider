package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oauth_provider/models"
)

var ACCESS_GROUP_MODEL = "access_groups"

func CreateAccessGroup(accessGroup *models.AccessGroup) (primitive.ObjectID, error) {
	accessGroup.ID = primitive.NewObjectID()
	return Create(accessGroup, ACCESS_GROUP_MODEL)
}

func UpdateAccessGroup(id primitive.ObjectID, accessGroup *models.AccessGroup) (primitive.ObjectID, error) {
	return Update(ACCESS_GROUP_MODEL, id, accessGroup)
}

func DeleteAccessGroup(id primitive.ObjectID) (*primitive.ObjectID, error) {
	return Delete[*models.AccessGroup](ACCESS_GROUP_MODEL, id)
}

func GetAccessGroups() ([]models.AccessGroup, error) {
	return GetList[models.AccessGroup](ACCESS_GROUP_MODEL)
}

func GetAccessGroup(id primitive.ObjectID) (models.AccessGroup, error) {
	return Get[models.AccessGroup](ACCESS_GROUP_MODEL, id)
}

func AccessGroupesFindByIdRessource(ressourceName string, ids []primitive.ObjectID) ([]models.AccessGroup, error) {
	attributes := map[string]interface{}{
		"_id": bson.M{
			"$in": ids,
		},
		"resources": bson.M{
			"$eq": ressourceName,
		},
	}

	return FindManyByAttributes[models.AccessGroup](ACCESS_GROUP_MODEL, attributes)
}
