package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type PostService struct {
	repo postRepository
}

type postRepository interface {
	Create(post models.Post) error
	GetAll() ([]models.Post, error)
	GetPublished() ([]models.Post, error)
	GetByID(id int64) (*models.Post, error)
	Update(id int64, post models.Post) error
	Delete(id int64) error
}

func NewPostService(repo postRepository) *PostService {
	return &PostService{repo: repo}
}

func validatePost(post models.Post) error {
	if strings.TrimSpace(post.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(post.Content) == "" {
		return errors.New("content is required")
	}
	if post.UniversityID != nil && strings.TrimSpace(*post.UniversityID) != "" {
		if _, err := uuid.Parse(strings.TrimSpace(*post.UniversityID)); err != nil {
			return errors.New("university_id must be a valid UUID")
		}
	}
	return nil
}

func (s *PostService) Create(post models.Post) error {
	if err := validatePost(post); err != nil {
		return err
	}
	return s.repo.Create(post)
}

func (s *PostService) GetAll() ([]models.Post, error) {
	return s.repo.GetAll()
}

func (s *PostService) GetPublished() ([]models.Post, error) {
	return s.repo.GetPublished()
}

func (s *PostService) GetByID(id int64) (*models.Post, error) {
	return s.repo.GetByID(id)
}

func (s *PostService) Update(id int64, post models.Post) error {
	if err := validatePost(post); err != nil {
		return err
	}
	return s.repo.Update(id, post)
}

func (s *PostService) Delete(id int64) error {
	return s.repo.Delete(id)
}
