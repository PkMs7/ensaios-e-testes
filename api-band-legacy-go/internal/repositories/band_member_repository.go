package repositories

import (
	"context"
	"database/sql"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/jmoiron/sqlx"
)

type BandMemberRepository struct {
	db *sqlx.DB
}

func NewBandMemberRepository(db *sqlx.DB) *BandMemberRepository {
	return &BandMemberRepository{db: db}
}

func (r *BandMemberRepository) CreateBandMember(ctx context.Context, bandMember *models.BandMember) error {
	query := `
		INSERT INTO band_member (band_id, artist_id, role_id, start_date, end_date)
		VALUES (:band_id, :artist_id, :role_id, :start_date, :end_date)
		RETURNING id, created_at
	`

	rows, err := r.db.NamedQueryContext(ctx, query, bandMember)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return rows.StructScan(bandMember)
	}

	return nil
}

func (r *BandMemberRepository) GetBandMembers(ctx context.Context) ([]models.BandMember, error) {
	query := `
		SELECT id,
			   band_id,
			   artist_id,
			   role_id,
			   start_date,
			   end_date,
			   created_at,
			   updated_at
		FROM band_member
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var bandMembers []models.BandMember

	for rows.Next() {
		var bandMember models.BandMember
		if err := rows.Scan(&bandMember.ID, &bandMember.BandId, &bandMember.ArtistId, &bandMember.RoleId, &bandMember.StartDate, &bandMember.EndDate, &bandMember.CreatedAt, &bandMember.UpdatedAt); err != nil {
			return nil, err
		}
		bandMembers = append(bandMembers, bandMember)
	}

	return bandMembers, nil
}

func (r *BandMemberRepository) GetBandMemberById(ctx context.Context, id int64) (*models.BandMember, error) {
	query := `
		SELECT id,
			   band_id,
			   artist_id,
			   role_id,
			   start_date,
			   end_date,
			   created_at,
			   updated_at
		FROM band_member
		WHERE id = $1
	`

	var bandMember models.BandMember
	err := r.db.GetContext(ctx, &bandMember, query, id)
	if err != nil {
		return nil, err
	}

	return &bandMember, nil
}

func (r *BandMemberRepository) UpdateBandMemberById(ctx context.Context, bandMember *models.BandMember) error {
	query := `
		UPDATE band_member
		SET start_date = :start_date,
			end_date = :end_date
		WHERE id = :id
	`

	result, err := r.db.NamedExecContext(ctx, query, bandMember)
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

func (r *BandMemberRepository) DeleteBandMemberById(ctx context.Context, id int64) error {
	query := `
		DELETE FROM band_member
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
