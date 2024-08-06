package services

import (
	"context"
	"github.com/weitien/admin/pkg/models"
	"github.com/weitien/admin/pkg/repositories"
)

type PermissionService interface {
	New(ctx context.Context, obj *models.Permission) error
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, obj *models.Permission) error
}

type permissionServiceImpl struct {
	repository repositories.PermissionRepository
}

func NewPermissionService() PermissionService {
	return &permissionServiceImpl{repositories.NewPermissionRepository()}
}

func (s *permissionServiceImpl) New(ctx context.Context, obj *models.Permission) error {
	return s.repository.New(ctx, repositories.GetDB(), obj)
}

func (s *permissionServiceImpl) Delete(ctx context.Context, id uint64) error {
	return s.repository.Delete(ctx, repositories.GetDB(), id)
}

func (s *permissionServiceImpl) Update(ctx context.Context, obj *models.Permission) error {
	return s.repository.Update(ctx, repositories.GetDB(), obj)
}
