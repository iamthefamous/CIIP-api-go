package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type UniversityService interface {
	GetAll() ([]models.University, error)
	GetPublished() ([]models.University, error)
	GetByID(id uuid.UUID) (*models.University, error)
	Create(university models.University) error
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, university models.University) error
}

type UniversityHandler struct{ service UniversityService }

func NewUniversityHandler(service UniversityService) *UniversityHandler {
	return &UniversityHandler{service: service}
}

func (h *UniversityHandler) GetAll(c *gin.Context) {
	universities, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, universities)
}

func (h *UniversityHandler) GetPublished(c *gin.Context) {
	universities, err := h.service.GetPublished()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, universities)
}

func (h *UniversityHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	university, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "university not found"})
		return
	}
	c.JSON(http.StatusOK, university)
}

func (h *UniversityHandler) Create(c *gin.Context) {
	var university models.University
	if err := c.ShouldBindJSON(&university); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(university); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "university created"})
}

func (h *UniversityHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err = h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "university deleted"})
}

func (h *UniversityHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var university models.University
	if err = c.ShouldBindJSON(&university); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = h.service.Update(id, university); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "university updated"})
}
