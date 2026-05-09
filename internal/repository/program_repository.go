package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamthefamous/CIIP-api-go/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProgramRepository struct {
	db *pgxpool.Pool
}

func NewProgramRepository(db *pgxpool.Pool) *ProgramRepository {
	return &ProgramRepository{db: db}
}

func (r *ProgramRepository) Create(program models.Program) error {
	query := `
		INSERT INTO programs (
			id,
			faculty_id,
			name,
			degree,
			duration_years,
			description,
			tuition_fee,
			is_published
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		uuid.New().String(),
		program.FacultyID,
		program.Name,
		program.Degree,
		program.DurationYears,
		program.Description,
		program.TuitionFee,
		program.IsPublished,
	)

	return err
}

func (r *ProgramRepository) GetAll() ([]models.Program, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, faculty_id, name, degree, duration_years, description, tuition_fee, is_published, created_at, updated_at FROM programs`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programs []models.Program
	for rows.Next() {
		var program models.Program
		err = rows.Scan(
			&program.ID,
			&program.FacultyID,
			&program.Name,
			&program.Degree,
			&program.DurationYears,
			&program.Description,
			&program.TuitionFee,
			&program.IsPublished,
			&program.CreatedAt,
			&program.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return programs, nil
}

func (r *ProgramRepository) GetByID(id uuid.UUID) (*models.Program, error) {
	query := `
		SELECT id, faculty_id, name, degree, duration_years, description, tuition_fee, is_published, created_at, updated_at
		FROM programs
		WHERE id = $1
	`

	var program models.Program
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&program.ID,
		&program.FacultyID,
		&program.Name,
		&program.Degree,
		&program.DurationYears,
		&program.Description,
		&program.TuitionFee,
		&program.IsPublished,
		&program.CreatedAt,
		&program.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &program, nil
}

func (r *ProgramRepository) Update(id uuid.UUID, program models.Program) error {
	query := `
		UPDATE programs
		SET
			faculty_id = $1,
			name = $2,
			degree = $3,
			duration_years = $4,
			description = $5,
			tuition_fee = $6,
			is_published = $7,
			updated_at = NOW()
		WHERE id = $8
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		program.FacultyID,
		program.Name,
		program.Degree,
		program.DurationYears,
		program.Description,
		program.TuitionFee,
		program.IsPublished,
		id,
	)

	return err
}

func (r *ProgramRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec(
		context.Background(),
		"DELETE FROM programs WHERE id = $1",
		id,
	)

	return err
}
