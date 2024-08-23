package repositories

import (
	"context"
	"github.com/weitien/admin/internal/models"
	"gorm.io/gorm"
)

type NavigationRepository interface {
	New(ctx context.Context, db *gorm.DB, nav *models.Navigation) error
	Delete(ctx context.Context, db *gorm.DB, id uint64) error
	Update(ctx context.Context, db *gorm.DB, nav *models.Navigation) error
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
