package models

import "time"

type Program struct {
	ID            string    `json:"id" db:"id"`
	FacultyID     string    `json:"faculty_id" db:"faculty_id"`
	Name          string    `json:"name" db:"name"`
	Degree        string    `json:"degree" db:"degree"`
	DurationYears int       `json:"duration_years" db:"duration_years"`
	Description   string    `json:"description" db:"description"`
	TuitionFee    float64   `json:"tuition_fee" db:"tuition_fee"`
	IsPublished   bool      `json:"is_published" db:"is_published"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
