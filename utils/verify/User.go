package verify

import (
	"log"

	"oauth_provider/models"
	"oauth_provider/db"
)

func User(username *string, password *string) (*models.User, error) {
	user, err := db.UserFindByUsernamePassword(username, password)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &user, nil
}