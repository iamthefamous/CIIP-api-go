package service

import (
	"errors"
	"testing"

	"github.com/iamthefamous/CIIP-api-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var errInvalidCreds = errors.New("invalid credentials")

type mockUserRepository struct {
	getByEmailFn func(email string) (*models.User, error)
}

func (m *mockUserRepository) GetByEmail(email string) (*models.User, error) {
	if m.getByEmailFn == nil {
		return nil, nil
	}
	return m.getByEmailFn(email)
}

func TestAuthServiceLoginSuccess(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	svc := NewAuthService(&mockUserRepository{
		getByEmailFn: func(email string) (*models.User, error) {
			if email != "admin@example.com" {
				t.Fatalf("expected email admin@example.com, got %s", email)
			}
			return &models.User{ID: "u1", Email: email, PasswordHash: string(hash), Role: "admin"}, nil
		},
	}, "test-secret")

	token, err := svc.Login("admin@example.com", "secret")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token == "" {
		t.Fatal("expected non-empty token")
	}
}

func TestAuthServiceLoginWrongPassword(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	svc := NewAuthService(&mockUserRepository{
		getByEmailFn: func(email string) (*models.User, error) {
			return &models.User{ID: "u1", Email: email, PasswordHash: string(hash), Role: "admin"}, nil
		},
	}, "test-secret")

	token, err := svc.Login("admin@example.com", "bad-password")
	if err == nil {
		t.Fatal("expected error")
	}
	if token != "" {
		t.Fatalf("expected empty token, got %s", token)
	}
}

func TestAuthServiceLoginRepositoryError(t *testing.T) {
	repoErr := errors.New("query failed")
	svc := NewAuthService(&mockUserRepository{
		getByEmailFn: func(email string) (*models.User, error) {
			return nil, repoErr
		},
	}, "test-secret")

	_, err := svc.Login("admin@example.com", "secret")
	if !errors.Is(err, repoErr) {
		t.Fatalf("expected %v, got %v", repoErr, err)
	}
}
