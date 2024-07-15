package repositories

import (
	"context"
	"gorm.io/gorm"

	"github.com/weitien/admin/models"
)

type AccountRepository interface {
	NewAccount(ctx context.Context, db *gorm.DB, acc *models.Account) error
	NewAccounts(ctx context.Context, db *gorm.DB, accounts []*models.Account) error
}

type accountRepositoryImpl struct{}

func NewAccountRepository() AccountRepository {
	return &accountRepositoryImpl{}
}

func (r *accountRepositoryImpl) NewAccount(ctx context.Context, db *gorm.DB, acc *models.Account) error {
	return db.WithContext(ctx).Create(acc).Error
}

func (r *accountRepositoryImpl) NewAccounts(ctx context.Context, db *gorm.DB, accounts []*models.Account) error {
	return db.WithContext(ctx).Create(accounts).Error
}
