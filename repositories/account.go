package repositories

import (
	"context"
	"gorm.io/gorm"

	"github.com/weitien/admin/models"
)

type AccountRepository interface {
	NewAccount(ctx context.Context, db *gorm.DB, acc *models.Account) error
	NewAccounts(ctx context.Context, db *gorm.DB, accounts []*models.Account) error
	DeleteAccounts(ctx context.Context, db *gorm.DB, userId uint64)
	ChangeAccountStatus(ctx context.Context, db *gorm.DB, id uint64, status uint)
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

func (r *accountRepositoryImpl) DeleteAccounts(ctx context.Context, db *gorm.DB, userId uint64) {
	sql := "DELETE FROM account WHERE user_id=?"
	db.Exec(sql, userId)
}

func (r *accountRepositoryImpl) ChangeAccountStatus(ctx context.Context, db *gorm.DB, id uint64, status uint) {
	db.Model(&models.Account{ID: id}).Update("status", status)
}
