package service

import (
	"errors"
	"testing"

	"github.com/iamthefamous/CIIP-api-go/internal/models"
)

type mockPostRepository struct {
	createFn       func(post models.Post) error
	getAllFn       func() ([]models.Post, error)
	getPublishedFn func() ([]models.Post, error)
	getByIDFn      func(id int64) (*models.Post, error)
	updateFn       func(id int64, post models.Post) error
	deleteFn       func(id int64) error
}

func (m *mockPostRepository) Create(post models.Post) error           { return m.createFn(post) }
func (m *mockPostRepository) GetAll() ([]models.Post, error)          { return m.getAllFn() }
func (m *mockPostRepository) GetPublished() ([]models.Post, error)    { return m.getPublishedFn() }
func (m *mockPostRepository) GetByID(id int64) (*models.Post, error)  { return m.getByIDFn(id) }
func (m *mockPostRepository) Update(id int64, post models.Post) error { return m.updateFn(id, post) }
func (m *mockPostRepository) Delete(id int64) error                   { return m.deleteFn(id) }

func TestPostServiceCreateValidatesUniversityID(t *testing.T) {
	repo := &mockPostRepository{createFn: func(post models.Post) error { return nil }}
	svc := NewPostService(repo)
	bad := "not-uuid"
	err := svc.Create(models.Post{Title: "t", Content: "c", UniversityID: &bad})
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestPostServiceGetPublished(t *testing.T) {
	expected := []models.Post{{ID: 1, Title: "a", Content: "b", IsPublished: true}}
	repo := &mockPostRepository{
		createFn:       func(post models.Post) error { return nil },
		getAllFn:       func() ([]models.Post, error) { return nil, nil },
		getPublishedFn: func() ([]models.Post, error) { return expected, nil },
		getByIDFn:      func(id int64) (*models.Post, error) { return nil, nil },
		updateFn:       func(id int64, post models.Post) error { return nil },
		deleteFn:       func(id int64) error { return nil },
	}
	svc := NewPostService(repo)
	posts, err := svc.GetPublished()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(posts) != 1 || !posts[0].IsPublished {
		t.Fatalf("unexpected posts: %+v", posts)
	}
}

func TestPostServiceUpdatePropagatesRepoError(t *testing.T) {
	repoErr := errors.New("db failed")
	repo := &mockPostRepository{
		createFn:       func(post models.Post) error { return nil },
		getAllFn:       func() ([]models.Post, error) { return nil, nil },
		getPublishedFn: func() ([]models.Post, error) { return nil, nil },
		getByIDFn:      func(id int64) (*models.Post, error) { return nil, nil },
		updateFn:       func(id int64, post models.Post) error { return repoErr },
		deleteFn:       func(id int64) error { return nil },
	}
	svc := NewPostService(repo)
	err := svc.Update(1, models.Post{Title: "t", Content: "c"})
	if !errors.Is(err, repoErr) {
		t.Fatalf("expected %v, got %v", repoErr, err)
	}
}
