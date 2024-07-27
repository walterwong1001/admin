package services

import (
	"context"
	"github.com/weitien/admin/models"
	"github.com/weitien/admin/repositories"
)

type RoleService interface {
	New(ctx context.Context, obj *models.Role) error
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, obj *models.Role) error
}

type roleServiceImpl struct {
	repository repositories.RoleRepository
}

func NewRoleService() RoleService {
	return &roleServiceImpl{repositories.NewRoleRepository()}
}

func (s *roleServiceImpl) New(ctx context.Context, obj *models.Role) error {
	return s.repository.New(ctx, repositories.GetDB(), obj)
}

func (s *roleServiceImpl) Delete(ctx context.Context, id uint64) error {
	return s.repository.Delete(ctx, repositories.GetDB(), id)
}

func (s *roleServiceImpl) Update(ctx context.Context, obj *models.Role) error {
	return s.repository.Update(ctx, repositories.GetDB(), obj)
}
