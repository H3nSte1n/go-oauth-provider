package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oauth_provider/models"
)

var SCOPE_MODEL = "scopes"

func CreateScope(scope *models.Scope) (primitive.ObjectID, error) {
	scope.ID = primitive.NewObjectID()
	return Create(scope, SCOPE_MODEL) //nolint:typecheck
}

func CreateScopes(scope []interface{}) ([]primitive.ObjectID, error) {
	return CreateMany(scope, SCOPE_MODEL) //nolint:typecheck
}

func GetScopes() ([]models.Scope, error) {
	return GetList[models.Scope](SCOPE_MODEL) //nolint:typecheck
}

func GetScope(id primitive.ObjectID) (models.Scope, error) {
	return Get[models.Scope](SCOPE_MODEL, id) //nolint:typecheck
}

func ScopeFindByNameIds(name string, ids []primitive.ObjectID) (models.Scope, error) {
	attributes := map[string]interface{}{
		"name": name,
		"_id": bson.M{
			"$in": ids,
		},
	}

	return FindByAttributes[models.Scope](SCOPE_MODEL, attributes) //nolint:typecheck
}
