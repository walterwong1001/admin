package repositories

import (
	"context"

	"github.com/weitien/admin/models"
)

type AccountRepository interface {
	NewAccount(ctx context.Context, acc *models.Account) error
	NewAccounts(ctx context.Context, accounts []*models.Account) error
}

type accountRepositoryImpl struct{}

func NewAccountRepository() AccountRepository {
	return &accountRepositoryImpl{}
}

func (r *accountRepositoryImpl) NewAccount(ctx context.Context, acc *models.Account) error {
	return DB.WithContext(ctx).Create(acc).Error
}

func (r *accountRepositoryImpl) NewAccounts(ctx context.Context, accounts []*models.Account) error {
	return DB.WithContext(ctx).Create(accounts).Error
}
