package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FacultyRepository struct {
	db *pgxpool.Pool
}

func NewFacultyRepository(db *pgxpool.Pool) *FacultyRepository {
	return &FacultyRepository{db: db}
}

func (r *FacultyRepository) Create(faculty models.Faculty) error {
	query := `
		INSERT INTO faculties (
			id,
			university_id,
			name,
			slug,
			description,
			is_published
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		uuid.New().String(),
		faculty.UniversityID,
		faculty.Name,
		faculty.Slug,
		faculty.Description,
		faculty.IsPublished,
	)

	return err
}

func (r *FacultyRepository) GetAll() ([]models.Faculty, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, university_id, name, slug, description, is_published, created_at, updated_at FROM faculties`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var faculties []models.Faculty
	for rows.Next() {
		var faculty models.Faculty
		err = rows.Scan(
			&faculty.ID,
			&faculty.UniversityID,
			&faculty.Name,
			&faculty.Slug,
			&faculty.Description,
			&faculty.IsPublished,
			&faculty.CreatedAt,
			&faculty.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		faculties = append(faculties, faculty)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return faculties, nil
}

func (r *FacultyRepository) GetByID(id uuid.UUID) (*models.Faculty, error) {
	query := `
		SELECT id, university_id, name, slug, description, is_published, created_at, updated_at
		FROM faculties
		WHERE id = $1
	`

	var faculty models.Faculty
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&faculty.ID,
		&faculty.UniversityID,
		&faculty.Name,
		&faculty.Slug,
		&faculty.Description,
		&faculty.IsPublished,
		&faculty.CreatedAt,
		&faculty.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &faculty, nil
}

func (r *FacultyRepository) Update(id uuid.UUID, faculty models.Faculty) error {
	query := `
		UPDATE faculties
		SET
			university_id = $1,
			name = $2,
			slug = $3,
			description = $4,
			is_published = $5,
			updated_at = NOW()
		WHERE id = $6
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		faculty.UniversityID,
		faculty.Name,
		faculty.Slug,
		faculty.Description,
		faculty.IsPublished,
		id,
	)

	return err
}

func (r *FacultyRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec(
		context.Background(),
		"DELETE FROM faculties WHERE id = $1",
		id,
	)

	return err
}
