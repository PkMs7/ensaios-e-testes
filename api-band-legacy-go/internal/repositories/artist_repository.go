package repositories

import (
	"context"
	"database/sql"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/jmoiron/sqlx"
)

type ArtistRepository struct {
	db *sqlx.DB
}

func NewArtistRepository(db *sqlx.DB) *ArtistRepository {
	return &ArtistRepository{db: db}
}

func (r *ArtistRepository) CreateArtist(ctx context.Context, artist *models.Artist) error {
	query := `
		INSERT INTO artist (name, birth_date, biography)
		VALUES (:name, :birth_date, :biography)
		RETURNING id, created_at
	`

	rows, err := r.db.NamedQueryContext(ctx, query, artist)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return rows.StructScan(artist)
	}

	return nil
}

func (r *ArtistRepository) GetArtists(ctx context.Context) ([]models.Artist, error) {
	query := `
		SELECT id,
			   name,
			   birth_date,
			   biography,
			   created_at,
			   updated_at
		FROM artist
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []models.Artist

	for rows.Next() {
		var artist models.Artist
		if err := rows.Scan(&artist.ID, &artist.Name, &artist.BirthDate, &artist.Biography, &artist.CreatedAt, &artist.UpdatedAt); err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}

	return artists, nil
}

func (r *ArtistRepository) GetArtistsById(ctx context.Context, id int64) (*models.Artist, error) {
	query := `
		SELECT id,
			   name,
			   birth_date,
			   biography,
			   created_at,
			   updated_at
		FROM artist
		WHERE id = $1
	`
	var artist models.Artist
	err := r.db.GetContext(ctx, &artist, query, id)
	if err != nil {
		return nil, err
	}

	return &artist, nil

}

func (r *ArtistRepository) UpdateArtistById(ctx context.Context, artist *models.Artist) error {
	query := `
		UPDATE artist
		SET name = :name,
			birth_date = :birth_date,
			biography = :biography
		WHERE id = :id
	`
	result, err := r.db.NamedExecContext(ctx, query, artist)
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

func (r *ArtistRepository) DeleteArtistById(ctx context.Context, id int64) error {
	query := `
		DELETE FROM artist
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
