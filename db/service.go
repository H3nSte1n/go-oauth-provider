package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"oauth_provider/models"
)

var SERVICE_MODEL = "services"

func CreateService(service *models.Service) (primitive.ObjectID, error) {
	service.ID = primitive.NewObjectID()
	return Create[*models.Service](service, SERVICE_MODEL)
}

func UpdateService(id primitive.ObjectID, service *models.Service) (primitive.ObjectID, error) {
	return Update[*models.Service](SERVICE_MODEL, id, service)
}

func GetServices() ([]models.Service, error) {
	return GetList[models.Service](SERVICE_MODEL)
}

func GetService(id primitive.ObjectID) (models.Service, error) {
	return Get[models.Service](SERVICE_MODEL, id)
}