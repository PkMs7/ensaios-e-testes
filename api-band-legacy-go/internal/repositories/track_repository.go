package repositories

import (
	"context"
	"database/sql"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/jmoiron/sqlx"
)

type TrackRepository struct {
	db *sqlx.DB
}

func NewTrackRepository(db *sqlx.DB) *TrackRepository {
	return &TrackRepository{db: db}
}

func (r *TrackRepository) CreateTrack(ctx context.Context, track *models.Track) error {
	query := `
		INSERT INTO track (album_id, title, track_number, duration_seconds)
		VALUES (:album_id, :title, :track_number, :duration_seconds)
		RETURNING id, created_at
	`

	rows, err := r.db.NamedQueryContext(ctx, query, track)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return rows.StructScan(track)
	}

	return nil
}

func (r *TrackRepository) GetTracks(ctx context.Context) ([]models.Track, error) {
	query := `
		SELECT id,
			   album_id,
			   title,
			   track_number,
			   duration_seconds,
			   created_at,
			   updated_at
		FROM track
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var tracks []models.Track

	for rows.Next() {
		var track models.Track
		if err := rows.Scan(&track.ID, &track.AlbumId, &track.Title, &track.TrackNumber, &track.DurationSeconds, &track.CreatedAt, &track.UpdatedAt); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetTrackById(ctx context.Context, id int64) (*models.Track, error) {
	query := `
		SELECT id,
			   album_id,
			   title,
			   track_number,
			   duration_seconds,
			   created_at,
			   updated_at
		FROM track
		WHERE id = $1
	`

	var track models.Track
	err := r.db.GetContext(ctx, &track, query, id)
	if err != nil {
		return nil, err
	}

	return &track, nil
}

func (r *TrackRepository) UpdateTrackById(ctx context.Context, track *models.Track) error {
	query := `
		UPDATE track
		SET album_id = :album_id,
			title = :title,
			track_number = :track_number,
			duration_seconds = :duration_seconds
		WHERE id = :id
	`

	result, err := r.db.NamedExecContext(ctx, query, track)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *TrackRepository) DeleteTrackById(ctx context.Context, id int64) error {
	query := `
		DELETE FROM track
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil

}
