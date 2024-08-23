package repositories

import (
	"context"
	"errors"

	"github.com/weitien/admin/internal/models"
	"gorm.io/gorm"
)

type AccessLogRepository interface {
	Append(ctx context.Context, db *gorm.DB, log *models.AccessLog) error
	Get(ctx context.Context, db *gorm.DB, id uint64) *models.AccessLog
}

type accessLogRepositoryImpl struct{}

func NewAccessLogRepository() AccessLogRepository {
	return &accessLogRepositoryImpl{}
}

func (a *accessLogRepositoryImpl) Append(ctx context.Context, db *gorm.DB, log *models.AccessLog) error {
	return db.WithContext(ctx).Create(log).Error
}

func (a *accessLogRepositoryImpl) Get(ctx context.Context, db *gorm.DB, id uint64) *models.AccessLog {
	var log models.AccessLog
	tx := db.WithContext(ctx).First(&log, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}
	return &log
}
