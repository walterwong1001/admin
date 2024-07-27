package repositories

import (
	"context"
	"github.com/weitien/admin/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	New(ctx context.Context, db *gorm.DB, obj *models.Role) error
	Delete(ctx context.Context, db *gorm.DB, id uint64) error
	Update(ctx context.Context, db *gorm.DB, obj *models.Role) error
}

type roleRepositoryImpl struct{}

func NewRoleRepository() RoleRepository {
	return &roleRepositoryImpl{}
}

func (r *roleRepositoryImpl) New(ctx context.Context, db *gorm.DB, obj *models.Role) error {
	return db.WithContext(ctx).Create(obj).Error
}

func (r *roleRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uint64) error {
	return db.WithContext(ctx).Delete(&models.Role{ID: id}).Error
}

func (r *roleRepositoryImpl) Update(ctx context.Context, db *gorm.DB, obj *models.Role) error {
	return db.WithContext(ctx).Updates(obj).Error
}
