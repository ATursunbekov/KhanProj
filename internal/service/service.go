package service

import (
	"github.com/ATursunbekov/KhanProj/internal/model"
	"github.com/ATursunbekov/KhanProj/internal/repository"
)

type Person interface {
	CreatePerson(p model.Person) error
	DeletePerson(id int) error
	UpdatePerson(p model.Person) error
	GetPersonByID(id int) (model.Person, error)
	GetAllPeople(filters map[string]string, limit, offset int) ([]model.Person, error)
}
type Service struct {
	Person
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Person: NewPersonService(&repo),
	}
}
