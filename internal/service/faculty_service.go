package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type facultyRepository interface {
	Create(faculty models.Faculty) error
	GetAll() ([]models.Faculty, error)
	GetPublished() ([]models.Faculty, error)
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

func validateFaculty(faculty models.Faculty) error {
	if strings.TrimSpace(faculty.UniversityID) == "" {
		return errors.New("university_id is required")
	}
	if _, err := uuid.Parse(strings.TrimSpace(faculty.UniversityID)); err != nil {
		return errors.New("university_id must be a valid UUID")
	}
	if strings.TrimSpace(faculty.Name) == "" {
		return errors.New("name is required")
	}
	return nil
}

func (s *FacultyService) Create(faculty models.Faculty) error {
	if err := validateFaculty(faculty); err != nil {
		return err
	}
	return s.repo.Create(faculty)
}

func (s *FacultyService) GetAll() ([]models.Faculty, error)       { return s.repo.GetAll() }
func (s *FacultyService) GetPublished() ([]models.Faculty, error) { return s.repo.GetPublished() }
func (s *FacultyService) GetByID(id uuid.UUID) (*models.Faculty, error) {
	return s.repo.GetByID(id)
}

func (s *FacultyService) Update(id uuid.UUID, faculty models.Faculty) error {
	if err := validateFaculty(faculty); err != nil {
		return err
	}
	return s.repo.Update(id, faculty)
}

func (s *FacultyService) Delete(id uuid.UUID) error { return s.repo.Delete(id) }
