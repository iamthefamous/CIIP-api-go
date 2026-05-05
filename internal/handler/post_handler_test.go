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
	createFn func(post models.Post) error
	getAllFn func() ([]models.Post, error)
}

func (m *mockPostService) CreatePost(post models.Post) error {
	if m.createFn == nil {
		return nil
	}
	return m.createFn(post)
}

func (m *mockPostService) GetAll() ([]models.Post, error) {
	if m.getAllFn == nil {
		return nil, nil
	}
	return m.getAllFn()
}

func TestPostHandlerCreateBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := NewPostHandler(&mockPostService{})
	r := gin.New()
	r.POST("/posts", handler.Create)

	req := httptest.NewRequest(http.MethodPost, "/posts", strings.NewReader("{bad"))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestPostHandlerCreateSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expected := models.Post{Title: "hello", Content: "world"}

	handler := NewPostHandler(&mockPostService{
		createFn: func(post models.Post) error {
			if post != expected {
				t.Fatalf("expected %+v, got %+v", expected, post)
			}
			return nil
		},
	})
	r := gin.New()
	r.POST("/posts", handler.Create)

	req := httptest.NewRequest(
		http.MethodPost,
		"/posts",
		strings.NewReader(`{"title":"hello","content":"world"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}
}

func TestPostHandlerCreateServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := NewPostHandler(&mockPostService{
		createFn: func(post models.Post) error {
			return errors.New("insert failed")
		},
	})
	r := gin.New()
	r.POST("/posts", handler.Create)

	req := httptest.NewRequest(
		http.MethodPost,
		"/posts",
		strings.NewReader(`{"title":"hello","content":"world"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

func TestPostHandlerGetAllSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	expected := []models.Post{
		{ID: 1, Title: "a", Content: "b"},
	}

	handler := NewPostHandler(&mockPostService{
		getAllFn: func() ([]models.Post, error) {
			return expected, nil
		},
	})
	r := gin.New()
	r.GET("/posts", handler.GetAll)

	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var got []models.Post
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if len(got) != len(expected) || got[0] != expected[0] {
		t.Fatalf("expected %+v, got %+v", expected, got)
	}
}

func TestPostHandlerGetAllServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := NewPostHandler(&mockPostService{
		getAllFn: func() ([]models.Post, error) {
			return nil, errors.New("query failed")
		},
	})
	r := gin.New()
	r.GET("/posts", handler.GetAll)

	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}
