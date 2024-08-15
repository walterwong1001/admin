package repositories

import (
	"context"

	"github.com/weitien/admin/pkg/models"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	Bind(ctx context.Context, db *gorm.DB, obj *models.UserRole) error
	UnBind(ctx context.Context, db *gorm.DB, obj *models.UserRole) error
	GetRolesByUser(ctx context.Context, db *gorm.DB, userId uint64) []uint64
}

type userRoleRepositoryImpl struct{}

func NewUserRoleRepository() UserRoleRepository {
	return &userRoleRepositoryImpl{}
}

func (r *userRoleRepositoryImpl) Bind(ctx context.Context, db *gorm.DB, obj *models.UserRole) error {
	return db.WithContext(ctx).Create(obj).Error
}

func (r *userRoleRepositoryImpl) UnBind(ctx context.Context, db *gorm.DB, obj *models.UserRole) error {
	return db.WithContext(ctx).Where("role_id=? AND user_id=?", obj.UserId, obj.RoleId).Delete(&models.UserRole{}).Error
}

func (r *userRoleRepositoryImpl) GetRolesByUser(ctx context.Context, db *gorm.DB, userId uint64) []uint64 {
	var roles []uint64
	db.Model(&models.UserRole{}).Where("user_id=?", userId).Pluck("role_id", &roles)
	return roles
}
