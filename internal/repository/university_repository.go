package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UniversityRepository struct {
	db *pgxpool.Pool
}

func NewUniversityRepository(db *pgxpool.Pool) *UniversityRepository {
	return &UniversityRepository{db: db}
}

func (r *UniversityRepository) Create(university models.University) error {
	query := `
		INSERT INTO universities (
			id,
			name,
			slug,
			description,
			address,
			phone,
			email,
			website,
			is_published
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		uuid.New().String(),
		university.Name,
		university.Slug,
		university.Description,
		university.Address,
		university.Phone,
		university.Email,
		university.Website,
		university.IsPublished,
	)

	return err
}

func (r *UniversityRepository) GetByID(id string) (*models.University, error) {
	var university models.University
	err := r.db.QueryRow(context.Background(),
		"SELECT * FROM universities WHERE id = $1",
		id).Scan(&university)

	if err != nil {
		return nil, err
	}
	return &university, nil
}
