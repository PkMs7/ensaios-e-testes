package services

import (
	"context"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/PkMs7/api-band-legacy-go/internal/repositories"
)

type TrackService struct {
	repo *repositories.TrackRepository
}

func NewTrackService(repo *repositories.TrackRepository) *TrackService {
	return &TrackService{repo: repo}
}

func (s *TrackService) CreateTrack(ctx context.Context, albumId int64, title string, trackNumber int64, durationSeconds int64) (*models.Track, error) {
	track := &models.Track{
		AlbumId:         albumId,
		Title:           title,
		TrackNumber:     trackNumber,
		DurationSeconds: durationSeconds,
	}

	err := s.repo.CreateTrack(ctx, track)

	return track, err
}

func (s *TrackService) GetTracks(ctx context.Context) ([]models.Track, error) {
	return s.repo.GetTracks(ctx)
}

func (s *TrackService) GetTrackById(ctx context.Context, id int64) (*models.Track, error) {
	return s.repo.GetTrackById(ctx, id)
}

func (s *TrackService) UpdateTrackById(ctx context.Context, track *models.Track) error {
	return s.repo.UpdateTrackById(ctx, track)
}

func (s *TrackService) DeleteTrackById(ctx context.Context, id int64) error {
	return s.repo.DeleteTrackById(ctx, id)
}
