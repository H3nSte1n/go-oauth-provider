package verify

import (
	"oauth_provider/models"
	"log"
	"oauth_provider/db"
	"oauth_provider/utils"
	"fmt"
)

func Client(client_secret string, client_id string, scope models.Scope, scope_err error) bool {
	credentials, credentials_err := checkIfCredentialsExist(client_secret, client_id)

	if credentials_err != nil || scope_err != nil {
		return false
	}

	if containsValue := utils.Contains(credentials.ScopeIDs, scope.ID); containsValue == false {		
		log.Print("Scope not found")
		return false
	}

	return true
}

func CheckIfScopeExists(ressource string) (*models.Scope, error) {
	fmt.Println("ressource: ", ressource)
	scopes, err := db.ScopesFindByName(ressource)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	if len(scopes) == 0 {
		log.Print("Scope not found")
		return nil, fmt.Errorf("Scope not found")
	}

	return &scopes[0], nil
}

func checkIfCredentialsExist(client_secret string, client_id string) (*models.Credential, error) {
	credentials, err := db.CredentialsFindByClientIdAndClientSecret(client_secret, client_id)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	fmt.Println("credentials: ", credentials)

	if len(credentials) == 0 {
		log.Print("Credentials not found")
		return nil, fmt.Errorf("Credentials not found")
	}

	return &credentials[0], nil
}