package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockAuthService struct {
	loginFn func(email, password string) (string, error)
}

func (m *mockAuthService) Login(email, password string) (string, error) {
	if m.loginFn == nil {
		return "", nil
	}
	return m.loginFn(email, password)
}

func TestAuthHandlerLoginBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewAuthHandler(&mockAuthService{})
	r := gin.New()
	r.POST("/admin/login", h.Login)

	req := httptest.NewRequest(http.MethodPost, "/admin/login", strings.NewReader("{bad"))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestAuthHandlerLoginUnauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewAuthHandler(&mockAuthService{
		loginFn: func(email, password string) (string, error) {
			return "", errors.New("invalid credentials")
		},
	})
	r := gin.New()
	r.POST("/admin/login", h.Login)

	req := httptest.NewRequest(http.MethodPost, "/admin/login", strings.NewReader(`{"email":"a@b.com","password":"x"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestAuthHandlerLoginSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewAuthHandler(&mockAuthService{
		loginFn: func(email, password string) (string, error) {
			if email != "admin@example.com" || password != "secret" {
				t.Fatalf("unexpected credentials: %s/%s", email, password)
			}
			return "jwt-token", nil
		},
	})
	r := gin.New()
	r.POST("/admin/login", h.Login)

	req := httptest.NewRequest(http.MethodPost, "/admin/login", strings.NewReader(`{"email":"admin@example.com","password":"secret"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "jwt-token") {
		t.Fatalf("expected token in response body, got %s", rec.Body.String())
	}
}
