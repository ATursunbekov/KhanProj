package handler

import (
	"github.com/ATursunbekov/KhanProj/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// @Summary Create new person
// @Tags Person
// @Accept json
// @Produce json
// @Param person body model.PersonInput true "Input for person creation"
// @Success 201 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /person/create [post]
func (h *Handler) Create(c *gin.Context) {
	var input model.PersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Errorf("Error binding json: %v", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	person := model.Person{
		Name:       input.Name,
		Surname:    input.Surname,
		Patronymic: input.Patronymic,
	}
	if err := h.service.CreatePerson(person); err != nil {
		logrus.Errorf("Error creating person: %v", err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.SuccessResponse{Status: "success"})
}

// @Summary Delete person by ID
// @Tags Person
// @Param id path int true "Person ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /person/delete/{id} [delete]
func (h *Handler) DeletePerson(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		logrus.Errorf("invalid id %s", idParam)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err = h.service.DeletePerson(id)
	if err != nil {
		logrus.Errorf("error deleting person %s", idParam)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{Status: "success"})
}

// @Summary Update person
// @Tags Person
// @Accept json
// @Produce json
// @Param person body model.Person true "Person info"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /person/update [put]
func (h *Handler) UpdatePerson(c *gin.Context) {
	var input model.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Errorf("invalid person %s", input.Name)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	if err := h.service.UpdatePerson(input); err != nil {
		logrus.Errorf("error updating person %s", input.Name)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{Status: "success"})
}

// @Summary Get person by ID
// @Tags Person
// @Param id path int true "Person ID"
// @Produce json
// @Success 200 {object} model.Person
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /person/getPerson/{id} [get]
func (h *Handler) GetPerson(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		logrus.Errorf("invalid id %s", idParam)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	res, err := h.service.GetPersonByID(id)
	if err != nil {
		logrus.Errorf("error getting person %s", idParam)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(200, res)
}

// @Summary Get all people with filters and pagination
// @Tags Person
// @Produce json
// @Param name query string false "Name filter"
// @Param gender query string false "Gender filter"
// @Param nationality query string false "Nationality filter"
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {array} model.Person
// @Failure 500 {object} model.ErrorResponse
// @Router /person/getAll [get]
func (h *Handler) GetAllPeople(c *gin.Context) {
	filters := map[string]string{}
	if name := c.Query("name"); name != "" {
		filters["name"] = name
	}
	if gender := c.Query("gender"); gender != "" {
		filters["gender"] = gender
	}
	if nationality := c.Query("nationality"); nationality != "" {
		filters["nationality"] = nationality
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit

	people, err := h.service.GetAllPeople(filters, limit, offset)
	if err != nil {
		logrus.Errorf("error getting people %s", filters)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, people)
}
