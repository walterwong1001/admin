package repositories

import (
	"context"
	"github.com/walterwong1001/admin/internal/models"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	New(ctx context.Context, db *gorm.DB, obj *models.Permission) error
	Delete(ctx context.Context, db *gorm.DB, id uint64) error
	Update(ctx context.Context, db *gorm.DB, obj *models.Permission) error
	All(ctx context.Context, db *gorm.DB) []*models.Permission
}

type permissionRepositoryImpl struct{}

func NewPermissionRepository() PermissionRepository {
	return &permissionRepositoryImpl{}
}

func (r *permissionRepositoryImpl) New(ctx context.Context, db *gorm.DB, obj *models.Permission) error {
	return db.WithContext(ctx).Create(obj).Error
}

func (r *permissionRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uint64) error {
	return db.WithContext(ctx).Delete(&models.Permission{ID: id}).Error
}

func (r *permissionRepositoryImpl) Update(ctx context.Context, db *gorm.DB, obj *models.Permission) error {
	return db.WithContext(ctx).Save(obj).Error
}

func (r *permissionRepositoryImpl) All(ctx context.Context, db *gorm.DB) []*models.Permission {
	var arr []*models.Permission
	db.WithContext(ctx).Find(&arr)
	return arr
}
