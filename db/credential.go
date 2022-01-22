package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"

	"oauth_provider/models"
)

var MODEL = "credentials"

func CreateCredential(credential *models.Credential) (primitive.ObjectID, error) {
	credential.ID = primitive.NewObjectID()
	return Create[*models.Credential](credential, MODEL)
}

func GetCredentials() ([]models.Credential, error) {
	return GetList[models.Credential](MODEL)
}

func GetCredential(id primitive.ObjectID) (models.Credential, error) {
	return Get[models.Credential](MODEL, id)
}

func CredentialsFindByClientIdAndClientSecret(client_secret string, client_id string) ([]models.Credential, error) {
	attributes := map[string]interface{}{
		"client_secret": client_secret,
		"client_id":     client_id,
	}

	fmt.Println("attributes", attributes)

	return FindByAttributes[models.Credential](MODEL, attributes)
}