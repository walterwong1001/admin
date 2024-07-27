package repositories

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/weitien/admin/models"
)

type AccountRepository interface {
	New(ctx context.Context, db *gorm.DB, acc *models.Account) error
	NewAccounts(ctx context.Context, db *gorm.DB, accounts []*models.Account) error
	Delete(ctx context.Context, db *gorm.DB, userId uint64) error
	ChangeStatus(ctx context.Context, db *gorm.DB, id uint64, status uint8) error
	GetByType(ctx context.Context, db *gorm.DB, identifier string, accType models.AccountType) *models.Account
}

type accountRepositoryImpl struct{}

func NewAccountRepository() AccountRepository {
	return &accountRepositoryImpl{}
}

func (r *accountRepositoryImpl) New(ctx context.Context, db *gorm.DB, acc *models.Account) error {
	return db.WithContext(ctx).Create(acc).Error
}

func (r *accountRepositoryImpl) NewAccounts(ctx context.Context, db *gorm.DB, accounts []*models.Account) error {
	return db.WithContext(ctx).Create(accounts).Error
}

func (r *accountRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, userId uint64) error {
	sql := "DELETE FROM account WHERE user_id=?"
	tx := db.WithContext(ctx).Exec(sql, userId)
	return tx.Error
}

func (r *accountRepositoryImpl) ChangeStatus(ctx context.Context, db *gorm.DB, id uint64, status uint8) error {
	return db.Model(&models.Account{ID: id}).Update("status", status).Error
}

func (r *accountRepositoryImpl) GetByType(ctx context.Context, db *gorm.DB, identifier string, accType models.AccountType) *models.Account {
	var acc models.Account
	tx := db.First(&acc, "identifier=? AND type=?", identifier, accType)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}
	return &acc
}
