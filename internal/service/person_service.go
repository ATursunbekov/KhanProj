package service

import (
	"encoding/json"
	"fmt"
	"github.com/ATursunbekov/KhanProj/internal/model"
	"github.com/ATursunbekov/KhanProj/internal/repository"
	"net/http"
)

type PersonService struct {
	repo *repository.Repository
}

func NewPersonService(r *repository.Repository) *PersonService {
	return &PersonService{repo: r}
}

func (s *PersonService) CreatePerson(p model.Person) error {
	age, gender, nationality := s.enrich(p.Name)
	p.Age = age
	p.Gender = gender
	p.Nationality = nationality
	return s.repo.Create(p)
}

func (s *PersonService) DeletePerson(id int) error {
	return s.repo.DeletePerson(id)
}

func (s *PersonService) UpdatePerson(p model.Person) error {
	return s.repo.UpdatePerson(p)
}

func (s *PersonService) GetPersonByID(id int) (model.Person, error) {
	return s.repo.GetPersonByID(id)
}

func (s *PersonService) GetAllPeople(filters map[string]string, limit, offset int) ([]model.Person, error) {
	return s.repo.GetAllPeople(filters, limit, offset)
}

func (s *PersonService) enrich(name string) (int, string, string) {
	var ag struct {
		Age int `json:"age"`
	}
	var gen struct {
		Gender string `json:"gender"`
	}
	var nat struct {
		Country []struct {
			CountryID string `json:"country_id"`
		} `json:"country"`
	}

	httpGetJson(fmt.Sprintf("https://api.agify.io/?name=%s", name), &ag)
	httpGetJson(fmt.Sprintf("https://api.genderize.io/?name=%s", name), &gen)
	httpGetJson(fmt.Sprintf("https://api.nationalize.io/?name=%s", name), &nat)

	nationality := ""
	if len(nat.Country) > 0 {
		nationality = nat.Country[0].CountryID
	}
	return ag.Age, gen.Gender, nationality
}

func httpGetJson(url string, target interface{}) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(target)
}
