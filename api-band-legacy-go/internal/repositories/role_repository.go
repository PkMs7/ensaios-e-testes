package repositories

import (
	"context"
	"database/sql"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/jmoiron/sqlx"
)

type RoleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Create(ctx context.Context, role *models.Role) error {
	query := `
			INSERT INTO role (name, description)
			VALUES (:name, :description)
			RETURNING id, created_at
	`

	rows, err := r.db.NamedQueryContext(ctx, query, role)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return rows.StructScan(role)
	}

	return nil
}

func (r *RoleRepository) GetRoles(ctx context.Context) ([]models.Role, error) {
	query := `
		SELECT id, 
			   name, 
			   description, 
			   created_at, 
			   updated_at
		FROM role	
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role

	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *RoleRepository) GetRoleById(ctx context.Context, id int64) (*models.Role, error) {
	query := `
		SELECT id, 
			   name, 
			   description, 
			   created_at, 
			   updated_at
		FROM role
		WHERE id = $1
	`

	var role models.Role
	err := r.db.GetContext(ctx, &role, query, id)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepository) UpdateRoleById(ctx context.Context, role *models.Role) error {
	query := `
		UPDATE role
		SET name = :name, 
			description = :description
		WHERE id = :id
	`

	result, err := r.db.NamedExecContext(ctx, query, role)
	if err != nil {
		return err
	}

	rowsAfected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (r *RoleRepository) DeleteRoleById(ctx context.Context, id int64) error {
	query := `
		DELETE FROM role
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
