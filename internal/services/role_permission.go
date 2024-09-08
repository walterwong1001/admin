package services

import (
	"context"

	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/repositories"
)

type RolePermissionService interface {
	Bind(ctx context.Context, obj *models.RolePermission) error
	UnBind(ctx context.Context, obj *models.RolePermission) error
	BatchBind(ctx context.Context, items []*models.RolePermission) error
	All(ctx context.Context) []*models.RolePermission
}

type rolePermissionServiceImpl struct {
	repository repositories.RolePermissionRepository
}

func NewRolePermissionService() RolePermissionService {
	return &rolePermissionServiceImpl{repository: repositories.NewRolePermissionRepository()}
}

func (s *rolePermissionServiceImpl) Bind(ctx context.Context, obj *models.RolePermission) error {
	return s.repository.Bind(ctx, repositories.GetDB(), obj)
}

func (s *rolePermissionServiceImpl) UnBind(ctx context.Context, obj *models.RolePermission) error {
	return s.repository.UnBind(ctx, repositories.GetDB(), obj)
}

func (s *rolePermissionServiceImpl) BatchBind(ctx context.Context, items []*models.RolePermission) error {
	return s.repository.BatchBind(ctx, repositories.GetDB(), items)
}

func (s *rolePermissionServiceImpl) All(ctx context.Context) []*models.RolePermission {
	return s.repository.All(ctx, repositories.GetDB())
}
