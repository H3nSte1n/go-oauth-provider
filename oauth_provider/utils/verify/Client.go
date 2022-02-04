package verify

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"oauth_provider/db"
	"oauth_provider/models"
	"oauth_provider/utils/array"
)

func Client(scopeId primitive.ObjectID, credentialScopeIds []primitive.ObjectID) bool {
	if containsValue := array.Contains[primitive.ObjectID](credentialScopeIds, scopeId); containsValue == false {
		log.Print("Scope not found")
		return false
	}

	return true
}

func ScopeExists(ressource string, scope_ids []primitive.ObjectID) (*models.Scope, error) {
	scopes, err := db.ScopeFindByNameIds(ressource, scope_ids)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &scopes, nil
}

func CredentialsExists(client_secret string, client_id string) (*models.Credential, error) {
	credentials, err := db.CredentialsFindByClientIdAndClientSecret(client_secret, client_id)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &credentials, nil
}
