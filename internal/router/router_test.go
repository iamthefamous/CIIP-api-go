package router

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/iamthefamous/CIIP-api-go/internal/handler"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type mockPostService struct{}

func (m *mockPostService) CreatePost(post models.Post) error {
	return nil
}

func (m *mockPostService) GetAll() ([]models.Post, error) {
	return []models.Post{}, nil
}

func TestNewRouterRegistersPostRoutes(t *testing.T) {
	postHandler := handler.NewPostHandler(&mockPostService{})
	r := NewRouter(postHandler)

	postReq := httptest.NewRequest(
		http.MethodPost,
		"/posts",
		strings.NewReader(`{"title":"hello","content":"world"}`),
	)
	postReq.Header.Set("Content-Type", "application/json")
	postRes := httptest.NewRecorder()
	r.ServeHTTP(postRes, postReq)

	if postRes.Code == http.StatusNotFound {
		t.Fatalf("expected POST /posts route to be registered")
	}

	getReq := httptest.NewRequest(http.MethodGet, "/posts", nil)
	getRes := httptest.NewRecorder()
	r.ServeHTTP(getRes, getReq)

	if getRes.Code == http.StatusNotFound {
		t.Fatalf("expected GET /posts route to be registered")
	}
}
