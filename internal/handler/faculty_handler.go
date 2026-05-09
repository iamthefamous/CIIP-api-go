package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type facultyService interface {
	GetAll() ([]models.Faculty, error)
	GetByID(id uuid.UUID) (*models.Faculty, error)
	Create(faculty models.Faculty) error
	Update(id uuid.UUID, faculty models.Faculty) error
	Delete(id uuid.UUID) error
}

type FacultyHandler struct {
	service facultyService
}

func NewFacultyHandler(service facultyService) *FacultyHandler {
	return &FacultyHandler{service: service}
}

func (h *FacultyHandler) GetAll(c *gin.Context) {
	faculties, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, faculties)
}

func (h *FacultyHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	faculty, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "faculty not found"})
		return
	}

	c.JSON(http.StatusOK, faculty)
}

func (h *FacultyHandler) Create(c *gin.Context) {
	var faculty models.Faculty
	if err := c.ShouldBindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(faculty); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "faculty created"})
}

func (h *FacultyHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var faculty models.Faculty
	if err = c.ShouldBindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.service.Update(id, faculty); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "faculty updated"})
}

func (h *FacultyHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err = h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "faculty deleted"})
}
