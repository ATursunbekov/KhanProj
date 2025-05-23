package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Configs struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SslMode  string
}

func NewPostgres(configs Configs) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		configs.Host, configs.Port, configs.Username, configs.DBName, configs.Password, configs.SslMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
