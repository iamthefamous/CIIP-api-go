package repository

import (
	"context"

	"github.com/iamthefamous/CIIP-api-go/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post models.Post) error {
	_, err := r.db.Exec(context.Background(),
		"INSERT INTO posts (title, context) VALUES ($1, $2)",
		post.Title, post.Content)

	return err
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	rows, err := r.db.Query(context.Background(),
		"SELECT id, title, content FROM posts")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		rows.Scan(&post.ID, &post.Title, &post.Content)
		posts = append(posts, post)
	}
	return posts, nil
}
