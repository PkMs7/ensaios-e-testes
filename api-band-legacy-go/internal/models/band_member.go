package models

import "time"

type BandMember struct {
	ID        int64     `db:"id" json:"id"`
	BandId    int64     `db:"band_id" json:"band_id"`
	ArtistId  int64     `db:"artist_id" json:"artist_id"`
	RoleId    int64     `db:"role_id" json:"role_id"`
	StartDate time.Time `db:"start_date" json:"start_date"`
	EndDate   time.Time `db:"end_date" json:"end_date"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
