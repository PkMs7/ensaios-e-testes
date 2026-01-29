package repositories

import (
	"context"
	"database/sql"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/jmoiron/sqlx"
)

type AlbumRepository struct {
	db *sqlx.DB
}

func NewAlbumRepository(db *sqlx.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (r *AlbumRepository) CreateAlbum(ctx context.Context, album *models.Album) error {
	query := `
		INSERT INTO album (band_id, title, release_date, album_type)
		VALUES (:band_id, :title, :release_date, :album_type)
		RETURNING id, created_at
	`

	rows, err := r.db.NamedQueryContext(ctx, query, album)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return rows.StructScan(album)
	}

	return nil
}

func (r *AlbumRepository) GetAlbums(ctx context.Context) ([]models.Album, error) {
	query := `
		SELECT id,
			   band_id,
			   title,
			   release_date,
			   album_type,
			   created_at,
			   updated_at
		FROM album
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var albums []models.Album

	for rows.Next() {
		var album models.Album
		if err := rows.Scan(&album.ID, &album.BandId, &album.Title, &album.ReleaseDate, &album.AlbumType, &album.CreatedAt, &album.UpdatedAt); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func (r *AlbumRepository) GetAlbumsById(ctx context.Context, id int64) (*models.Album, error) {
	query := `
		SELECT id,
			   band_id,
			   title,
			   release_date,
			   album_type,
			   created_at,
			   updated_at
		FROM album
		WHERE id = $1
	`

	var album models.Album
	err := r.db.GetContext(ctx, &album, query, id)
	if err != nil {
		return nil, err
	}

	return &album, nil

}

func (r *AlbumRepository) UpdateAlbumById(ctx context.Context, album *models.Album) error {
	query := `
		UPDATE album
		SET band_id = :band_id,
			title = :title,
			release_date = :release_date,
			album_type = :album_type
		WHERE id = :id
	`

	result, err := r.db.NamedExecContext(ctx, query, album)
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

func (r *AlbumRepository) DeleteAlbumById(ctx context.Context, id int64) error {
	query := `
		DELETE FROM album
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
