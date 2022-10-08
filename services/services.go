package services

import (
	"github.com/person/CRUD-api/models"
	"github.com/person/CRUD-api/utils"
)

type Services struct {
	PersonService *PersonService
}

func InitServices(config utils.AppConfig) *Services {

	template := &PersonService{make(models.PersonList), 0}
	return &Services{
		PersonService: template,
	}

}
