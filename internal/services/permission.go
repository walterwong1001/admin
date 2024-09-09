package services

import (
	"context"
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/repositories"
	"github.com/walterwong1001/gin_common_libs/page"
)

type PermissionService interface {
	page.PaginationHelper[*models.Permission, *models.PermissionFilter]
	New(ctx context.Context, obj *models.Permission) error
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, obj *models.Permission) error
	All(ctx context.Context) []*models.Permission
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

func (s *permissionServiceImpl) All(ctx context.Context) []*models.Permission {
	return s.repository.All(ctx, repositories.GetDB())
}

func (s *permissionServiceImpl) Pagination(ctx context.Context, p page.Paginator[*models.Permission], filter *models.PermissionFilter) error {
	//TODO implement me
	return nil
}
