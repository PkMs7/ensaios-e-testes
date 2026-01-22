package services

import (
	"context"
	"errors"
	"time"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/PkMs7/api-band-legacy-go/internal/repositories"
)

type ArtistService struct {
	repo *repositories.ArtistRepository
}

func NewArtistRepository(repo *repositories.ArtistRepository) *ArtistService {
	return &ArtistService{repo: repo}
}

func (s *ArtistService) CreateArtist(ctx context.Context, name string, birth_date time.Time, biography string) (*models.Artist, error) {
	if name == "" {
		return nil, errors.New("Artist name is required")
	}

	artist := &models.Artist{
		Name:      name,
		BirthDate: birth_date,
		Biography: biography,
	}
	err := s.repo.CreateArtist(ctx, artist)

	return artist, err

}

func (s *ArtistService) GetArtists(ctx context.Context) ([]models.Artist, error) {
	return s.repo.GetArtists(ctx)
}

func (s *ArtistService) GetArtistsById(ctx context.Context, id int64) (*models.Artist, error) {
	return s.repo.GetArtistsById(ctx, id)
}

func (s *ArtistService) UpdateArtistById(ctx context.Context, artist *models.Artist) error {
	return s.repo.UpdateArtistById(ctx, artist)
}

func (s *ArtistService) DeleteArtistById(ctx context.Context, id int64) error {
	return s.repo.DeleteArtistById(ctx, id)
}
