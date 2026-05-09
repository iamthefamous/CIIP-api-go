package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type mockFacultyRepository struct{}

func (m *mockFacultyRepository) Create(faculty models.Faculty) error { return nil }
func (m *mockFacultyRepository) GetAll() ([]models.Faculty, error)   { return []models.Faculty{}, nil }
func (m *mockFacultyRepository) GetPublished() ([]models.Faculty, error) {
	return []models.Faculty{}, nil
}
func (m *mockFacultyRepository) GetByID(id uuid.UUID) (*models.Faculty, error) {
	return &models.Faculty{}, nil
}
func (m *mockFacultyRepository) Update(id uuid.UUID, faculty models.Faculty) error { return nil }
func (m *mockFacultyRepository) Delete(id uuid.UUID) error                         { return nil }

func TestFacultyServiceCreateInvalidUniversityID(t *testing.T) {
	svc := NewFacultyService(&mockFacultyRepository{})
	err := svc.Create(models.Faculty{UniversityID: "bad-id", Name: "Faculty"})
	if err == nil {
		t.Fatal("expected validation error")
	}
}
