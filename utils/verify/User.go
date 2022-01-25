package verify

import (
	"golang.org/x/crypto/bcrypt"
	"log"

	"oauth_provider/models"
	"oauth_provider/db"
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