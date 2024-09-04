package services

import (
	"context"

	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/repositories"
)

type RoleService interface {
	New(ctx context.Context, obj *models.Role) error
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, obj *models.Role) error
	All(ctx context.Context) []*models.Role
	Pagination(ctx context.Context, p page.Paginator[*models.Role], filter *models.RoleFilter) error
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

func (s *roleServiceImpl) All(ctx context.Context) []*models.Role {
	return s.repository.All(ctx, repositories.GetDB())
}

func (s *roleServiceImpl) Pagination(ctx context.Context, p page.Paginator[*models.Role], filter *models.RoleFilter) error {
	return s.repository.Pagination(ctx, repositories.GetDB(), p, filter)
}
