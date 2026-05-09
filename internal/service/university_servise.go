package service

import (
	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type universityRepository interface {
	Create(university models.University) error
	GetAll() ([]models.University, error)
	GetPublished() ([]models.University, error)
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (*models.University, error)
	Update(id uuid.UUID, university models.University) error
}

type UniversityService struct {
	repo universityRepository
}

func NewUniversityService(repo universityRepository) *UniversityService {
	return &UniversityService{repo: repo}
}

func (s *UniversityService) Create(university models.University) error {
	return s.repo.Create(university)
}
func (s *UniversityService) GetAll() ([]models.University, error)       { return s.repo.GetAll() }
func (s *UniversityService) GetPublished() ([]models.University, error) { return s.repo.GetPublished() }
func (s *UniversityService) Delete(id uuid.UUID) error                  { return s.repo.Delete(id) }
func (s *UniversityService) GetByID(id uuid.UUID) (*models.University, error) {
	return s.repo.GetByID(id)
}
func (s *UniversityService) Update(id uuid.UUID, university models.University) error {
	return s.repo.Update(id, university)
}
