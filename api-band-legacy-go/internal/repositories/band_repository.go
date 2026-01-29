package repositories

import (
	"context"
	"database/sql"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/jmoiron/sqlx"
)

type BandRepository struct {
	db *sqlx.DB
}

func NewBandRepository(db *sqlx.DB) *BandRepository {
	return &BandRepository{db: db}
}

func (r *BandRepository) CreateBand(ctx context.Context, band *models.Band) error {
	query := `
		INSERT INTO band (name, formed_date, ended_date, description, is_active)
		VALUES (:name, :formed_date, :ended_date, :description, :is_active)
		RETURNING id, created_at
	`

	rows, err := r.db.NamedQueryContext(ctx, query, band)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return rows.StructScan(band)
	}

	return nil
}

func (r *BandRepository) GetBands(ctx context.Context) ([]models.Band, error) {
	query := `
		SELECT id,
			   name,
			   formed_date,
			   ended_date,
			   description,
			   is_active,
			   created_at,
			   updated_at
		FROM band
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var bands []models.Band

	for rows.Next() {
		var band models.Band
		if err := rows.Scan(&band.ID, &band.Name, &band.FormedDate, &band.EndedDate, &band.Description, &band.IsActive, &band.CreatedAt, &band.UpdatedAt); err != nil {
			return nil, err
		}
		bands = append(bands, band)
	}

	return bands, nil

}

func (r *BandRepository) GetBandById(ctx context.Context, id int64) (*models.Band, error) {
	query := `
		SELECT id,
			   name,
			   formed_date,
			   ended_date,
			   description,
			   is_active,
			   created_at,
			   updated_at
		FROM band
		WHERE id = $1
	`

	var band models.Band
	err := r.db.GetContext(ctx, &band, query, id)
	if err != nil {
		return nil, err
	}

	return &band, nil
}

func (r *BandRepository) UpdateBandById(ctx context.Context, band *models.Band) error {
	query := `
		UPDATE band
		SET name = :name,
			formed_date = :formed_date,
			ended_date = :ended_date,
			description = :description,
			is_active = :is_active
		WHERE id = :id
	`

	result, err := r.db.NamedExecContext(ctx, query, band)
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

func (r *BandRepository) DeleteBandById(ctx context.Context, id int64) error {
	query := `
		DELETE FROM band
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
