package models

import "time"

type Faculty struct {
	ID           string    `json:"id" db:"id"`
	UniversityID string    `json:"university_id" db:"university_id"`
	Name         string    `json:"name" db:"name"`
	Slug         string    `json:"slug" db:"slug"`
	Description  string    `json:"description" db:"description"`
	IsPublished  bool      `json:"is_published" db:"is_published"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
