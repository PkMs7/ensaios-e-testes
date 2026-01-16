package services

import (
	"context"
	"errors"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/PkMs7/api-band-legacy-go/internal/repositories"
)

type RoleService struct {
	repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) CreateRole(ctx context.Context, name string, description string) (*models.Role, error) {
	if name == "" {
		return nil, errors.New("Role name is required")
	}

	if description == "" {
		return nil, errors.New("Role description is required")
	}

	role := &models.Role{
		Name:        name,
		Description: description,
	}
	err := s.repo.Create(ctx, role)

	return role, err
}

func (s *RoleService) GetRoles(ctx context.Context) ([]models.Role, error) {
	return s.repo.GetRoles(ctx)
}

func (s *RoleService) GetRoleById(ctx context.Context, id int64) (*models.Role, error) {
	return s.repo.GetRoleById(ctx, id)
}

func (s *RoleService) UpdateRoleById(ctx context.Context, role *models.Role) error {
	return s.repo.UpdateRoleById(ctx, role)
}

func (s *RoleService) DeleteRoleById(ctx context.Context, id int64) error {
	return s.repo.DeleteRoleById(ctx, id)
}
