package repositories

import (
	"context"
	"errors"

	"github.com/walterwong1001/admin/internal/models"
	"gorm.io/gorm"
)

type NavigationRepository interface {
	New(ctx context.Context, db *gorm.DB, nav *models.Navigation) error
	Delete(ctx context.Context, db *gorm.DB, id uint64) error
	Update(ctx context.Context, db *gorm.DB, nav *models.Navigation) error
	All(ctx context.Context, db *gorm.DB) []*models.Navigation
}

type navigationRepositoryImpl struct{}

func NewNavigationRepository() NavigationRepository {
	return &navigationRepositoryImpl{}
}

func (r *navigationRepositoryImpl) New(ctx context.Context, db *gorm.DB, nav *models.Navigation) error {
	return db.WithContext(ctx).Create(nav).Error
}

func (r *navigationRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uint64) error {
	return db.WithContext(ctx).Delete(&models.Navigation{ID: id}).Error
}

func (r *navigationRepositoryImpl) Update(ctx context.Context, db *gorm.DB, nav *models.Navigation) error {
	return db.WithContext(ctx).Updates(nav).Error
}

func (r *navigationRepositoryImpl) All(ctx context.Context, db *gorm.DB) []*models.Navigation {
	var navs []*models.Navigation
	tx := db.WithContext(ctx).Find(&navs)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}
	return navs
}
