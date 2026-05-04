package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type PostHandler struct {
	service postService
}

type postService interface {
	CreatePost(post models.Post) error
	GetAll() ([]models.Post, error)
}

func NewPostHandler(service postService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) Create(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post created"})
}

func (h *PostHandler) GetAll(c *gin.Context) {
	posts, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}
