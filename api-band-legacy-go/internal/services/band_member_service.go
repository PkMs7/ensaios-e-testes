package services

import (
	"context"

	"github.com/PkMs7/api-band-legacy-go/internal/models"
	"github.com/PkMs7/api-band-legacy-go/internal/repositories"
)

type BandMemberService struct {
	repo *repositories.BandMemberRepository
}

func NewBandMemberService(repo *repositories.BandMemberRepository) *BandMemberService {
	return &BandMemberService{repo: repo}
}

func (s *BandMemberService) CreateBandMember(ctx context.Context, bandMember *models.BandMember) (*models.BandMember, error) {
	err := s.repo.CreateBandMember(ctx, bandMember)

	return bandMember, err
}

func (s *BandMemberService) GetBandMembers(ctx context.Context) ([]models.BandMember, error) {
	return s.repo.GetBandMembers(ctx)
}

func (s *BandMemberService) GetBandMemberById(ctx context.Context, id int64) (*models.BandMember, error) {
	return s.repo.GetBandMemberById(ctx, id)
}

func (s *BandMemberService) UpdateBandMemberById(ctx context.Context, bandMember *models.BandMember) error {
	return s.repo.UpdateBandMemberById(ctx, bandMember)
}

func (s *BandMemberService) DeleteBandmemberById(ctx context.Context, id int64) error {
	return s.repo.DeleteBandMemberById(ctx, id)
}
