package service

import (
	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type programRepository interface {
	Create(program models.Program) error
	GetAll() ([]models.Program, error)
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

func (s *ProgramService) Create(program models.Program) error {
	return s.repo.Create(program)
}

func (s *ProgramService) GetAll() ([]models.Program, error) {
	return s.repo.GetAll()
}

func (s *ProgramService) GetByID(id uuid.UUID) (*models.Program, error) {
	return s.repo.GetByID(id)
}

func (s *ProgramService) Update(id uuid.UUID, program models.Program) error {
	return s.repo.Update(id, program)
}

func (s *ProgramService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
