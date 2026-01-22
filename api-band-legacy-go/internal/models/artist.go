package models

import "time"

type Artist struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
	Biography string    `db:"biography" json:"biography"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
