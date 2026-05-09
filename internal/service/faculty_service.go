package service

import (
	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type facultyRepository interface {
	Create(faculty models.Faculty) error
	GetAll() ([]models.Faculty, error)
	GetByID(id uuid.UUID) (*models.Faculty, error)
	Update(id uuid.UUID, faculty models.Faculty) error
	Delete(id uuid.UUID) error
}

type FacultyService struct {
	repo facultyRepository
}

func NewFacultyService(repo facultyRepository) *FacultyService {
	return &FacultyService{repo: repo}
}

func (s *FacultyService) Create(faculty models.Faculty) error {
	return s.repo.Create(faculty)
}

func (s *FacultyService) GetAll() ([]models.Faculty, error) {
	return s.repo.GetAll()
}

func (s *FacultyService) GetByID(id uuid.UUID) (*models.Faculty, error) {
	return s.repo.GetByID(id)
}

func (s *FacultyService) Update(id uuid.UUID, faculty models.Faculty) error {
	return s.repo.Update(id, faculty)
}

func (s *FacultyService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
