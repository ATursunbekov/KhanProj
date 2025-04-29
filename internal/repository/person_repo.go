package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ATursunbekov/KhanProj/internal/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PersonRepo struct {
	db *sqlx.DB
}

func NewPersonRepo(db *sqlx.DB) *PersonRepo {
	return &PersonRepo{db: db}
}

func (r *PersonRepo) Create(p model.Person) error {
	query := `
        INSERT INTO persons (name, surname, patronymic, age, gender, nationality)
        VALUES (:name, :surname, :patronymic, :age, :gender, :nationality)
        RETURNING id;
    `
	rows, err := r.db.NamedQuery(query, p)
	if err != nil {
		logrus.Errorf("Error inserting person into database: %v", err)
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&p.ID)
		if err != nil {
			logrus.Errorf("Error inserting person into database: %v", err)
			return err
		}
	}

	return nil
}

func (r *PersonRepo) DeletePerson(id int) error {
	query := `DELETE FROM persons WHERE id = $1`

	res, err := r.db.Exec(query, id)
	if err != nil {
		logrus.Errorf("Error deleting person into database: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No person found with id %d", id)
	}

	return nil
}

func (r *PersonRepo) UpdatePerson(p model.Person) error {
	query := `
		UPDATE persons 
		SET name = :name, surname = :surname, patronymic = :patronymic,
		    age = :age, gender = :gender, nationality = :nationality
		WHERE id = :id
	`

	res, err := r.db.NamedExec(query, p)
	if err != nil {
		logrus.Errorf("Error updating person into database: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.Errorf("Error updating person into database: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no person found with id %d", p.ID)
	}

	return nil
}

func (r *PersonRepo) GetPersonByID(id int) (model.Person, error) {
	var person model.Person

	query := `SELECT * FROM persons WHERE id = $1`
	err := r.db.Get(&person, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logrus.Errorf("No person found with id %d", id)
			return person, fmt.Errorf("person with id %d not found", id)
		}
		logrus.Errorf("Error getting person from database: %v", err)
		return person, err
	}

	return person, nil
}

func (r *PersonRepo) GetAllPeople(filters map[string]string, limit, offset int) ([]model.Person, error) {
	people := []model.Person{}

	query := `SELECT * FROM persons WHERE 1=1`
	args := []interface{}{}
	argIdx := 1

	for key, value := range filters {
		query += fmt.Sprintf(" AND %s = $%d", key, argIdx)
		args = append(args, value)
		argIdx++
	}

	query += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argIdx, argIdx+1)
	args = append(args, limit, offset)

	err := r.db.Select(&people, query, args...)
	return people, err
}
