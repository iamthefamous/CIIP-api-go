package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type mockProgramRepository struct{}

func (m *mockProgramRepository) Create(program models.Program) error { return nil }
func (m *mockProgramRepository) GetAll() ([]models.Program, error)   { return []models.Program{}, nil }
func (m *mockProgramRepository) GetPublished() ([]models.Program, error) {
	return []models.Program{}, nil
}
func (m *mockProgramRepository) GetByID(id uuid.UUID) (*models.Program, error) {
	return &models.Program{}, nil
}
func (m *mockProgramRepository) Update(id uuid.UUID, program models.Program) error { return nil }
func (m *mockProgramRepository) Delete(id uuid.UUID) error                         { return nil }

func TestProgramServiceCreateInvalidFacultyID(t *testing.T) {
	svc := NewProgramService(&mockProgramRepository{})
	err := svc.Create(models.Program{FacultyID: "bad-id", Name: "Program", Degree: "BSc"})
	if err == nil {
		t.Fatal("expected validation error")
	}
}
