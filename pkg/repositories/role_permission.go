package repositories

import (
	"context"

	"github.com/weitien/admin/pkg/models"
	"gorm.io/gorm"
)

type RolePermissionRepository interface {
	Bind(ctx context.Context, db *gorm.DB, obj *models.RolePermission) error
	UnBind(ctx context.Context, db *gorm.DB, obj *models.RolePermission) error
	BatchBind(ctx context.Context, db *gorm.DB, items []*models.RolePermission) error
	All(ctx context.Context, db *gorm.DB) []*models.RolePermission
}

func NewRolePermissionRepository() RolePermissionRepository {
	return &rolePermissionRepositoryImpl{}
}

type rolePermissionRepositoryImpl struct{}

func (r *rolePermissionRepositoryImpl) Bind(ctx context.Context, db *gorm.DB, obj *models.RolePermission) error {
	return db.WithContext(ctx).Create(obj).Error
}

func (r *rolePermissionRepositoryImpl) UnBind(ctx context.Context, db *gorm.DB, obj *models.RolePermission) error {
	return db.WithContext(ctx).Where("role_id=? AND permission_id=?", obj.RoleId, obj.PermissionId).Delete(&models.RolePermission{}).Error
}

func (r *rolePermissionRepositoryImpl) BatchBind(ctx context.Context, db *gorm.DB, items []*models.RolePermission) error {
	return db.WithContext(ctx).CreateInBatches(items, 100).Error
}

func (r *rolePermissionRepositoryImpl) All(ctx context.Context, db *gorm.DB) []*models.RolePermission {
	var items []*models.RolePermission
	db.WithContext(ctx).Find(&items)
	return items
}
