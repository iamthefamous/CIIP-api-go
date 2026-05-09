package router

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/handler"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type mockPostService struct{}

func (m *mockPostService) Create(post models.Post) error        { return nil }
func (m *mockPostService) GetAll() ([]models.Post, error)       { return []models.Post{}, nil }
func (m *mockPostService) GetPublished() ([]models.Post, error) { return []models.Post{}, nil }
func (m *mockPostService) GetByID(id int64) (*models.Post, error) {
	p := &models.Post{ID: id}
	return p, nil
}
func (m *mockPostService) Update(id int64, post models.Post) error { return nil }
func (m *mockPostService) Delete(id int64) error                   { return nil }

type mockAuthService struct {
	loginFn func(email, password string) (string, error)
}

func (m *mockAuthService) Login(email, password string) (string, error) {
	if m.loginFn == nil {
		return "", nil
	}
	return m.loginFn(email, password)
}

type mockUniversityService struct{}

func (m *mockUniversityService) GetAll() ([]models.University, error) {
	return []models.University{}, nil
}
func (m *mockUniversityService) GetPublished() ([]models.University, error) {
	return []models.University{}, nil
}
func (m *mockUniversityService) GetByID(id uuid.UUID) (*models.University, error) {
	u := &models.University{ID: id.String()}
	return u, nil
}
func (m *mockUniversityService) Create(university models.University) error               { return nil }
func (m *mockUniversityService) Delete(id uuid.UUID) error                               { return nil }
func (m *mockUniversityService) Update(id uuid.UUID, university models.University) error { return nil }

type mockFacultyService struct{}

func (m *mockFacultyService) GetAll() ([]models.Faculty, error)       { return []models.Faculty{}, nil }
func (m *mockFacultyService) GetPublished() ([]models.Faculty, error) { return []models.Faculty{}, nil }
func (m *mockFacultyService) GetByID(id uuid.UUID) (*models.Faculty, error) {
	f := &models.Faculty{ID: id.String()}
	return f, nil
}
func (m *mockFacultyService) Create(faculty models.Faculty) error               { return nil }
func (m *mockFacultyService) Update(id uuid.UUID, faculty models.Faculty) error { return nil }
func (m *mockFacultyService) Delete(id uuid.UUID) error                         { return nil }

type mockProgramService struct{}

func (m *mockProgramService) GetAll() ([]models.Program, error)       { return []models.Program{}, nil }
func (m *mockProgramService) GetPublished() ([]models.Program, error) { return []models.Program{}, nil }
func (m *mockProgramService) GetByID(id uuid.UUID) (*models.Program, error) {
	p := &models.Program{ID: id.String()}
	return p, nil
}
func (m *mockProgramService) Create(program models.Program) error               { return nil }
func (m *mockProgramService) Update(id uuid.UUID, program models.Program) error { return nil }
func (m *mockProgramService) Delete(id uuid.UUID) error                         { return nil }

func TestNewRouterRegistersPublicRoutes(t *testing.T) {
	postHandler := handler.NewPostHandler(&mockPostService{})
	authHandler := handler.NewAuthHandler(&mockAuthService{loginFn: func(email, password string) (string, error) { return "token", nil }})
	universityHandler := handler.NewUniversityHandler(&mockUniversityService{})
	facultyHandler := handler.NewFacultyHandler(&mockFacultyService{})
	programHandler := handler.NewProgramHandler(&mockProgramService{})

	r := NewRouter(postHandler, authHandler, universityHandler, facultyHandler, programHandler, "secret", "", nil)

	getReq := httptest.NewRequest(http.MethodGet, "/posts", nil)
	getRes := httptest.NewRecorder()
	r.ServeHTTP(getRes, getReq)
	if getRes.Code == http.StatusNotFound {
		t.Fatalf("expected GET /posts route to be registered")
	}

	getByIDReq := httptest.NewRequest(http.MethodGet, "/posts/1", nil)
	getByIDRes := httptest.NewRecorder()
	r.ServeHTTP(getByIDRes, getByIDReq)
	if getByIDRes.Code == http.StatusNotFound {
		t.Fatalf("expected GET /posts/:id route to be registered")
	}

	loginReq := httptest.NewRequest(http.MethodPost, "/admin/login", strings.NewReader(`{"email":"admin@example.com","password":"secret"}`))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRes := httptest.NewRecorder()
	r.ServeHTTP(loginRes, loginReq)
	if loginRes.Code == http.StatusNotFound {
		t.Fatalf("expected POST /admin/login route to be registered")
	}
}

func TestNewRouterProtectsAdminRoutes(t *testing.T) {
	postHandler := handler.NewPostHandler(&mockPostService{})
	authHandler := handler.NewAuthHandler(&mockAuthService{loginFn: func(email, password string) (string, error) { return "", errors.New("invalid credentials") }})
	universityHandler := handler.NewUniversityHandler(&mockUniversityService{})
	facultyHandler := handler.NewFacultyHandler(&mockFacultyService{})
	programHandler := handler.NewProgramHandler(&mockProgramService{})

	r := NewRouter(postHandler, authHandler, universityHandler, facultyHandler, programHandler, "secret", "", nil)

	req := httptest.NewRequest(http.MethodGet, "/admin/posts", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	if res.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, res.Code)
	}
}
