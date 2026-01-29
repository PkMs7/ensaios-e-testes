package services

import (
	"context"
	"time"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/PkMs7/api-band-legacy-go/internal/repositories"
)

type AlbumService struct {
	repo *repositories.AlbumRepository
}

func NewAlbumService(repo *repositories.AlbumRepository) *AlbumService {
	return &AlbumService{repo: repo}
}

func (s *AlbumService) CreateAlbum(ctx context.Context, band_id int64, title string, releaseDate time.Time, albumType string) (*models.Album, error) {
	album := &models.Album{
		BandId:      band_id,
		Title:       title,
		ReleaseDate: releaseDate,
		AlbumType:   albumType,
	}
	err := s.repo.CreateAlbum(ctx, album)

	return album, err
}

func (s *AlbumService) GetAlbums(ctx context.Context) ([]models.Album, error) {
	return s.repo.GetAlbums(ctx)
}

func (s *AlbumService) GetAlbumsById(ctx context.Context, id int64) (*models.Album, error) {
	return s.repo.GetAlbumsById(ctx, id)
}

func (s *AlbumService) UpdateAlbumById(ctx context.Context, album *models.Album) error {
	return s.repo.UpdateAlbumById(ctx, album)
}

func (s *AlbumService) DeleteAlbumById(ctx context.Context, id int64) error {
	return s.repo.DeleteAlbumById(ctx, id)
}
