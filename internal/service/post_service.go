package service

import (
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type PostService struct {
	repo postRepository
}

type postRepository interface {
	Create(post models.Post) error
	GetAll() ([]models.Post, error)
}

func NewPostService(repo postRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post models.Post) error {
	return s.repo.Create(post)
}

func (s *PostService) GetAll() ([]models.Post, error) {
	return s.repo.GetAll()
}
