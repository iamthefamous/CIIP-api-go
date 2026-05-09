package models

import "time"

type Post struct {
	ID           int64     `json:"id"`
	UniversityID *string   `json:"university_id,omitempty"`
	Title        string    `json:"title" binding:"required"`
	Content      string    `json:"content" binding:"required"`
	IsPublished  bool      `json:"is_published"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
