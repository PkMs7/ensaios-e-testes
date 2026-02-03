package models

import "time"

type Track struct {
	ID              int64     `db:"id" json:"id"`
	AlbumId         int64     `db:"album_id" json:"album_id"`
	Title           string    `db:"title" json:"title"`
	TrackNumber     int64     `db:"track_number" json:"track_number"`
	DurationSeconds int64     `db:"duration_seconds" json:"duration_seconds"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}
