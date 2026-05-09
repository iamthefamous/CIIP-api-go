package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type programRepository interface {
	Create(program models.Program) error
	GetAll() ([]models.Program, error)
	GetPublished() ([]models.Program, error)
	GetByID(id uuid.UUID) (*models.Program, error)
	Update(id uuid.UUID, program models.Program) error
	Delete(id uuid.UUID) error
}

type ProgramService struct {
	repo programRepository
}

func NewProgramService(repo programRepository) *ProgramService {
	return &ProgramService{repo: repo}
}

func validateProgram(program models.Program) error {
	if strings.TrimSpace(program.FacultyID) == "" {
		return errors.New("faculty_id is required")
	}
	if _, err := uuid.Parse(strings.TrimSpace(program.FacultyID)); err != nil {
		return errors.New("faculty_id must be a valid UUID")
	}
	if strings.TrimSpace(program.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(program.Degree) == "" {
		return errors.New("degree is required")
	}
	return nil
}

func (s *ProgramService) Create(program models.Program) error {
	if err := validateProgram(program); err != nil {
		return err
	}
	return s.repo.Create(program)
}

func (s *ProgramService) GetAll() ([]models.Program, error)       { return s.repo.GetAll() }
func (s *ProgramService) GetPublished() ([]models.Program, error) { return s.repo.GetPublished() }
func (s *ProgramService) GetByID(id uuid.UUID) (*models.Program, error) {
	return s.repo.GetByID(id)
}

func (s *ProgramService) Update(id uuid.UUID, program models.Program) error {
	if err := validateProgram(program); err != nil {
		return err
	}
	return s.repo.Update(id, program)
}

func (s *ProgramService) Delete(id uuid.UUID) error { return s.repo.Delete(id) }
