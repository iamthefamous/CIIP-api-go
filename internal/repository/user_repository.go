package repository

import (
	"context"

	"github.com/iamthefamous/CIIP-api-go/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := r.db.QueryRow(
		context.Background(),
		`SELECT user_id, email, password_hash, role
		 FROM users
		 WHERE email = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetProfileRoleByID(userID string) (string, error) {
	var role string

	err := r.db.QueryRow(
		context.Background(),
		`SELECT role
		 FROM profiles
		 WHERE id = $1`,
		userID,
	).Scan(&role)
	if err != nil {
		return "", err
	}

	return role, nil
}
