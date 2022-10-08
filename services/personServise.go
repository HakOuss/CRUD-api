package services

import (
	"github.com/person/CRUD-api/models"
)

type PersonService struct {
	person models.PersonList
	index  int
}

func (service *PersonService) Create(person *models.Person) {
	service.person[service.index] = person
	service.index++
}

func (service *PersonService) List() models.PersonList {
	return service.person
}

func (service *PersonService) Update(index int, person *models.Person) {
	service.person[index] = person
}
func (service *PersonService) Delete(index int) {
	delete(service.person, index)
}
