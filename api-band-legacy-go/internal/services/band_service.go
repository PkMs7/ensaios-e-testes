package services

import (
	"context"
	"errors"
	"time"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/PkMs7/api-band-legacy-go/internal/repositories"
)

type BandService struct {
	repo *repositories.BandRepository
}

func NewBandService(repo *repositories.BandRepository) *BandService {
	return &BandService{repo: repo}
}

func (s *BandService) CreateBand(ctx context.Context, name string, formed_date time.Time, ended_date time.Time, description string, is_active bool) (*models.Band, error) {
	if name == "" {
		return nil, errors.New("Band name is required")
	}

	band := &models.Band{
		Name:        name,
		FormedDate:  formed_date,
		EndedDate:   ended_date,
		Description: description,
		IsActive:    is_active,
	}
	err := s.repo.CreateBand(ctx, band)

	return band, err

}

func (s *BandService) GetBands(ctx context.Context) ([]models.Band, error) {
	return s.repo.GetBands(ctx)
}

func (s *BandService) GetBandById(ctx context.Context, id int64) (*models.Band, error) {
	return s.repo.GetBandById(ctx, id)
}

func (s *BandService) UpdateBandById(ctx context.Context, band *models.Band) error {
	return s.repo.UpdateBandById(ctx, band)
}

func (s *BandService) DeleteBandById(ctx context.Context, id int64) error {
	return s.repo.DeleteBandById(ctx, id)
}
