package models

import "time"

type Album struct {
	ID          int64     `db:"id" json:"id"`
	BandId      int64     `db:"band_id" json:"band_id"`
	Title       string    `db:"title" json:"title"`
	ReleaseDate time.Time `db:"release_date" json:"release_date"`
	AlbumType   string    `db:"album_type" json:"album_type"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
