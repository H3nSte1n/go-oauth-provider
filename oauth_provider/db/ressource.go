package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oauth_provider/models"
)

var RESSOURCE_MODEL = "ressources"

func CreateRessource(ressource *models.Ressource) (primitive.ObjectID, error) {
	ressource.ID = primitive.NewObjectID()
	return Create(ressource, RESSOURCE_MODEL) //nolint:typecheck
}

func UpdateRessource(id primitive.ObjectID, ressource *models.Ressource) (primitive.ObjectID, error) {
	return Update(RESSOURCE_MODEL, id, ressource) //nolint:typecheck
}

func GetRessources() ([]models.Ressource, error) {
	return GetList[models.Ressource](RESSOURCE_MODEL) //nolint:typecheck
}

func GetRessource(id primitive.ObjectID) (models.Ressource, error) {
	return Get[models.Ressource](RESSOURCE_MODEL, id) //nolint:typecheck
}
