package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oauth_provider/models"
)

var SCOPE_MODEL = "scopes"

func CreateScope(scope *models.Scope) (primitive.ObjectID, error) {
	scope.ID = primitive.NewObjectID()
	return Create[*models.Scope](scope, SCOPE_MODEL)
}

func CreateScopes(scope []interface{}) ([]primitive.ObjectID, error) {
	return CreateMany(scope, SCOPE_MODEL)
}

func GetScopes() ([]models.Scope, error) {
	return GetList[models.Scope](SCOPE_MODEL)
}

func GetScope(id primitive.ObjectID) (models.Scope, error) {
	return Get[models.Scope](SCOPE_MODEL, id)
}

func ScopesFindByName(name string) ([]models.Scope, error) {
	attributes := map[string]interface{}{
		"name": name,
	}

	return FindByAttributes[models.Scope](SCOPE_MODEL, attributes)
}