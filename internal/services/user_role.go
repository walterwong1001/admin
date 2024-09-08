package services

import (
	"context"

	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/repositories"
)

type UserRoleService interface {
	Bind(ctx context.Context, obj *models.UserRole) error
	UnBind(ctx context.Context, obj *models.UserRole) error
	GetRolesByUser(ctx context.Context, userId uint64) []uint64
}

type userRoleServiceImpl struct {
	repository repositories.UserRoleRepository
}

func NewUserRoleService() UserRoleService {
	return &userRoleServiceImpl{repositories.NewUserRoleRepository()}
}

func (s *userRoleServiceImpl) Bind(ctx context.Context, obj *models.UserRole) error {
	return s.repository.Bind(ctx, repositories.GetDB(), obj)
}

func (s *userRoleServiceImpl) UnBind(ctx context.Context, obj *models.UserRole) error {
	return s.repository.UnBind(ctx, repositories.GetDB(), obj)
}

func (s *userRoleServiceImpl) GetRolesByUser(ctx context.Context, userId uint64) []uint64 {
	return s.repository.GetRolesByUser(ctx, repositories.GetDB(), userId)
}
