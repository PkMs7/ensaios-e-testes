package models

import "time"

type Band struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	FormedDate  time.Time `db:"formed_date" json:"formed_date"`
	EndedDate   time.Time `db:"ended_date" json:"ended_date"`
	Description string    `db:"description" json:"description"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
