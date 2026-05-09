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
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
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

func (r *UniversityRepository) GetByID(id uuid.UUID) (*models.University, error) {
	query := `
		SELECT id, name, slug, description, address, phone, email, website, is_published, created_at, updated_at
		FROM universities
		WHERE id = $1
	`

	var university models.University
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&university.ID,
		&university.Name,
		&university.Slug,
		&university.Description,
		&university.Address,
		&university.Phone,
		&university.Email,
		&university.Website,
		&university.IsPublished,
		&university.CreatedAt,
		&university.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &university, nil
}

func (r *UniversityRepository) Update(id uuid.UUID, university models.University) error {
	query := `
		UPDATE universities
		SET
			name = $1,
			slug = $2,
			description = $3,
			address = $4,
			phone = $5,
			email = $6,
			website = $7,
			is_published = $8,
			updated_at = NOW()
		WHERE id = $9
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		university.Name,
		university.Slug,
		university.Description,
		university.Address,
		university.Phone,
		university.Email,
		university.Website,
		university.IsPublished,
		id,
	)

	return err
}

func (r *UniversityRepository) GetAll() ([]models.University, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, name, slug, description, address, phone, email, website, is_published, created_at, updated_at FROM universities`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var universities []models.University
	for rows.Next() {
		var university models.University
		err = rows.Scan(
			&university.ID,
			&university.Name,
			&university.Slug,
			&university.Description,
			&university.Address,
			&university.Phone,
			&university.Email,
			&university.Website,
			&university.IsPublished,
			&university.CreatedAt,
			&university.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		universities = append(universities, university)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return universities, nil
}

func (r *UniversityRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec(
		context.Background(),
		"DELETE FROM universities WHERE id = $1",
		id,
	)

	return err
}
