package verify

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"oauth_provider/db"
	"oauth_provider/models"
)

func User(username *string, password *string) (*models.User, error) {
	user, err := db.UserFindByUsername(username)
	validPassword := validPassword(password, &user.Password)

	if err != nil || !validPassword {
		log.Print(err)
		return nil, err
	}

	return &user, nil
}

func validPassword(password *string, hash *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	return err == nil
}
