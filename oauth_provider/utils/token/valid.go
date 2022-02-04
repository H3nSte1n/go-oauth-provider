package token

import (
	"fmt"
	"oauth_provider/db"
)

func Valid(ressourceName string, tokenString string) error {
	claims, err := verify(tokenString)
	if err != nil {
		return err
	}

	accessGroups, err := db.AccessGroupesFindByIdRessource(ressourceName, claims.OwnClaims.AccessGroups)
	if err != nil || len(accessGroups) == 0 {
		return fmt.Errorf("You are not authorized to access this ressource")
	}

	return nil
}
