package handler

import (
	_ "github.com/ATursunbekov/KhanProj/docs"
	"github.com/ATursunbekov/KhanProj/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (s *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	person := router.Group("/person")
	{
		person.POST("/create", s.Create)
		person.DELETE("/delete/:id", s.DeletePerson)
		person.GET("/getPerson/:id", s.GetPerson)
		person.PUT("/update", s.UpdatePerson)
		person.GET("/getAll/", s.GetAllPeople)
	}

	return router
}
