package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type mockPostService struct {
	createFn       func(post models.Post) error
	getAllFn       func() ([]models.Post, error)
	getPublishedFn func() ([]models.Post, error)
	getByIDFn      func(id int64) (*models.Post, error)
	updateFn       func(id int64, post models.Post) error
	deleteFn       func(id int64) error
}

func (m *mockPostService) Create(post models.Post) error           { return m.createFn(post) }
func (m *mockPostService) GetAll() ([]models.Post, error)          { return m.getAllFn() }
func (m *mockPostService) GetPublished() ([]models.Post, error)    { return m.getPublishedFn() }
func (m *mockPostService) GetByID(id int64) (*models.Post, error)  { return m.getByIDFn(id) }
func (m *mockPostService) Update(id int64, post models.Post) error { return m.updateFn(id, post) }
func (m *mockPostService) Delete(id int64) error                   { return m.deleteFn(id) }

func TestPostHandlerCreateBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewPostHandler(&mockPostService{
		createFn:       func(post models.Post) error { return nil },
		getAllFn:       func() ([]models.Post, error) { return nil, nil },
		getPublishedFn: func() ([]models.Post, error) { return nil, nil },
		getByIDFn:      func(id int64) (*models.Post, error) { return nil, nil },
		updateFn:       func(id int64, post models.Post) error { return nil },
		deleteFn:       func(id int64) error { return nil },
	})
	r := gin.New()
	r.POST("/posts", handler.Create)

	req := httptest.NewRequest(http.MethodPost, "/posts", strings.NewReader("{bad"))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestPostHandlerGetAllPublicSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expected := []models.Post{{ID: 1, Title: "a", Content: "b", IsPublished: true}}
	handler := NewPostHandler(&mockPostService{
		createFn:       func(post models.Post) error { return nil },
		getAllFn:       func() ([]models.Post, error) { return nil, nil },
		getPublishedFn: func() ([]models.Post, error) { return expected, nil },
		getByIDFn:      func(id int64) (*models.Post, error) { return nil, nil },
		updateFn:       func(id int64, post models.Post) error { return nil },
		deleteFn:       func(id int64) error { return nil },
	})
	r := gin.New()
	r.GET("/posts", handler.GetAllPublic)

	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, rec.Code)
	}
	var got []models.Post
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(got) != 1 || !got[0].IsPublished {
		t.Fatalf("unexpected response: %+v", got)
	}
}

func TestPostHandlerUpdateServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewPostHandler(&mockPostService{
		createFn:       func(post models.Post) error { return nil },
		getAllFn:       func() ([]models.Post, error) { return nil, nil },
		getPublishedFn: func() ([]models.Post, error) { return nil, nil },
		getByIDFn:      func(id int64) (*models.Post, error) { return nil, nil },
		updateFn:       func(id int64, post models.Post) error { return errors.New("bad") },
		deleteFn:       func(id int64) error { return nil },
	})
	r := gin.New()
	r.PUT("/posts/:id", handler.Update)

	req := httptest.NewRequest(http.MethodPut, "/posts/1", strings.NewReader(`{"title":"hello","content":"world"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, rec.Code)
	}
}
