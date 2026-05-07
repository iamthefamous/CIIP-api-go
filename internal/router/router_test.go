package router

import (
	"errors"
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

type mockAuthService struct {
	loginFn func(email, password string) (string, error)
}

func (m *mockAuthService) Login(email, password string) (string, error) {
	if m.loginFn == nil {
		return "", nil
	}
	return m.loginFn(email, password)
}

func TestNewRouterRegistersPublicRoutes(t *testing.T) {
	postHandler := handler.NewPostHandler(&mockPostService{})
	authHandler := handler.NewAuthHandler(&mockAuthService{
		loginFn: func(email, password string) (string, error) {
			return "token", nil
		},
	})
	r := NewRouter(postHandler, authHandler, "secret", "", nil)

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

	loginReq := httptest.NewRequest(
		http.MethodPost,
		"/admin/login",
		strings.NewReader(`{"email":"admin@example.com","password":"secret"}`),
	)
	loginReq.Header.Set("Content-Type", "application/json")
	loginRes := httptest.NewRecorder()
	r.ServeHTTP(loginRes, loginReq)
	if loginRes.Code == http.StatusNotFound {
		t.Fatalf("expected POST /admin/login route to be registered")
	}
}

func TestNewRouterProtectsAdminRoutes(t *testing.T) {
	postHandler := handler.NewPostHandler(&mockPostService{})
	authHandler := handler.NewAuthHandler(&mockAuthService{
		loginFn: func(email, password string) (string, error) {
			return "", errors.New("invalid credentials")
		},
	})
	r := NewRouter(postHandler, authHandler, "secret", "", nil)

	req := httptest.NewRequest(http.MethodGet, "/admin/posts", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, res.Code)
	}
}
