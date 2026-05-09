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
	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO posts (university_id, title, content, is_published) VALUES ($1, $2, $3, $4)`,
		post.UniversityID,
		post.Title,
		post.Content,
		post.IsPublished,
	)

	return err
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, university_id, title, content, is_published, created_at, updated_at FROM posts ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.ID,
			&post.UniversityID,
			&post.Title,
			&post.Content,
			&post.IsPublished,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) GetPublished() ([]models.Post, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, university_id, title, content, is_published, created_at, updated_at FROM posts WHERE is_published = true ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.ID,
			&post.UniversityID,
			&post.Title,
			&post.Content,
			&post.IsPublished,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) GetByID(id int64) (*models.Post, error) {
	var post models.Post
	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, university_id, title, content, is_published, created_at, updated_at FROM posts WHERE id = $1`,
		id,
	).Scan(
		&post.ID,
		&post.UniversityID,
		&post.Title,
		&post.Content,
		&post.IsPublished,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Update(id int64, post models.Post) error {
	_, err := r.db.Exec(
		context.Background(),
		`UPDATE posts SET university_id=$1, title=$2, content=$3, is_published=$4, updated_at=NOW() WHERE id=$5`,
		post.UniversityID,
		post.Title,
		post.Content,
		post.IsPublished,
		id,
	)

	return err
}

func (r *PostRepository) Delete(id int64) error {
	_, err := r.db.Exec(
		context.Background(),
		`DELETE FROM posts WHERE id = $1`,
		id,
	)
	return err
}
