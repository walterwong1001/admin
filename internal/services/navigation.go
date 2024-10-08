package services

import (
	"context"
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/repositories"
)

type NavigationServicer interface {
	New(ctx context.Context, nav *models.Navigation) error
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, nav *models.Navigation) error
	All(ctx context.Context) []*models.Navigation
}

type navigationService struct {
	repository repositories.NavigationRepository
}

func NewNavigationService() NavigationServicer {
	return &navigationService{repositories.NewNavigationRepository()}
}

func (s *navigationService) New(ctx context.Context, nav *models.Navigation) error {
	return s.repository.New(ctx, repositories.GetDB(), nav)
}

func (s *navigationService) Delete(ctx context.Context, id uint64) error {
	return s.repository.Delete(ctx, repositories.GetDB(), id)
}

func (s *navigationService) Update(ctx context.Context, nav *models.Navigation) error {
	return s.repository.Update(ctx, repositories.GetDB(), nav)
}

func (s *navigationService) All(ctx context.Context) []*models.Navigation {
	return s.repository.All(ctx, repositories.GetDB())
}
