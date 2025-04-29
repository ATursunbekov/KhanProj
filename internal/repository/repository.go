package repository

import (
	"github.com/ATursunbekov/KhanProj/internal/model"
	"github.com/jmoiron/sqlx"
)

type Person interface {
	Create(p model.Person) error
	DeletePerson(id int) error
	UpdatePerson(p model.Person) error
	GetPersonByID(id int) (model.Person, error)
	GetAllPeople(filters map[string]string, limit, offset int) ([]model.Person, error)
}

type Repository struct {
	Person
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Person: NewPersonRepo(db),
	}
}
