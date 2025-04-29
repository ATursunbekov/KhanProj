package main

import (
	"github.com/ATursunbekov/KhanProj/configs"
	_ "github.com/ATursunbekov/KhanProj/docs"
	"github.com/ATursunbekov/KhanProj/internal/handler"
	"github.com/ATursunbekov/KhanProj/internal/repository"
	"github.com/ATursunbekov/KhanProj/internal/server"
	"github.com/ATursunbekov/KhanProj/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// @title KhanProj API
// @version 1.0
// @description API for managing people
// @host localhost:8080
// @BasePath /
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	configs.LoadEnv()

	db, err := repository.NewPostgres(repository.Configs{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Errorf("failed to connect to database: %v", err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(*repo)
	handler := handler.New(service)

	srv := new(server.Server)
	if err := srv.Start(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
