package service

import (
	"errors"
	"testing"

	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type mockPostRepository struct {
	createFn func(post models.Post) error
	getAllFn func() ([]models.Post, error)
}

func (m *mockPostRepository) Create(post models.Post) error {
	if m.createFn == nil {
		return nil
	}
	return m.createFn(post)
}

func (m *mockPostRepository) GetAll() ([]models.Post, error) {
	if m.getAllFn == nil {
		return nil, nil
	}
	return m.getAllFn()
}

func TestPostServiceCreatePost(t *testing.T) {
	expected := models.Post{Title: "hello", Content: "world"}
	called := false

	repo := &mockPostRepository{
		createFn: func(post models.Post) error {
			called = true
			if post != expected {
				t.Fatalf("expected %+v, got %+v", expected, post)
			}
			return nil
		},
	}
	svc := NewPostService(repo)

	if err := svc.CreatePost(expected); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !called {
		t.Fatal("expected repository Create to be called")
	}
}

func TestPostServiceCreatePostError(t *testing.T) {
	repoErr := errors.New("db failed")
	repo := &mockPostRepository{
		createFn: func(post models.Post) error {
			return repoErr
		},
	}
	svc := NewPostService(repo)

	err := svc.CreatePost(models.Post{Title: "hello"})
	if !errors.Is(err, repoErr) {
		t.Fatalf("expected %v, got %v", repoErr, err)
	}
}

func TestPostServiceGetAll(t *testing.T) {
	expected := []models.Post{
		{ID: 1, Title: "a", Content: "b"},
		{ID: 2, Title: "c", Content: "d"},
	}

	repo := &mockPostRepository{
		getAllFn: func() ([]models.Post, error) {
			return expected, nil
		},
	}
	svc := NewPostService(repo)

	posts, err := svc.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(posts) != len(expected) {
		t.Fatalf("expected %d posts, got %d", len(expected), len(posts))
	}
	for i := range expected {
		if posts[i] != expected[i] {
			t.Fatalf("expected %+v, got %+v", expected[i], posts[i])
		}
	}
}

func TestPostServiceGetAllError(t *testing.T) {
	repoErr := errors.New("query failed")
	repo := &mockPostRepository{
		getAllFn: func() ([]models.Post, error) {
			return nil, repoErr
		},
	}
	svc := NewPostService(repo)

	posts, err := svc.GetAll()
	if !errors.Is(err, repoErr) {
		t.Fatalf("expected %v, got %v", repoErr, err)
	}
	if posts != nil {
		t.Fatalf("expected nil posts, got %+v", posts)
	}
}
