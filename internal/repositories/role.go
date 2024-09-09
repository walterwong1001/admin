package repositories

import (
	"context"
	"errors"
	"github.com/walterwong1001/gin_common_libs/page"

	"github.com/walterwong1001/admin/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	New(ctx context.Context, db *gorm.DB, obj *models.Role) error
	Delete(ctx context.Context, db *gorm.DB, id uint64) error
	Update(ctx context.Context, db *gorm.DB, obj *models.Role) error
	All(ctx context.Context, db *gorm.DB) []*models.Role
	Pagination(ctx context.Context, db *gorm.DB, p page.Paginator[*models.Role], filter *models.RoleFilter) error
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
	return db.WithContext(ctx).Save(obj).Error
}

func (r *roleRepositoryImpl) All(ctx context.Context, db *gorm.DB) []*models.Role {
	var roles []*models.Role
	tx := db.WithContext(ctx).Find(&roles)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}
	return roles
}

func (r *roleRepositoryImpl) Pagination(ctx context.Context, db *gorm.DB, p page.Paginator[*models.Role], filter *models.RoleFilter) error {

	query := db.WithContext(ctx)
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.Start != 0 {
		query = query.Where("create_time >= ?", filter.Start)
	}
	if filter.End != 0 {
		query = query.Where("create_time <= ?", filter.End)
	}

	var total int64
	err := query.Model(&models.Role{}).Count(&total).Error
	if err != nil {
		return err
	}
	p.SetTotal(total)

	var roles []*models.Role

	err = query.Limit(p.GetPageSize()).Offset(p.Offset()).Find(&roles).Error
	if err != nil {
		return err
	}
	p.SetItems(roles)
	return nil
}
