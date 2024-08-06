package repositories

import (
	"context"
	"errors"
	"github.com/weitien/admin/pkg/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	New(cxt context.Context, db *gorm.DB, user *models.User) error
	Get(cxt context.Context, id uint64) *models.User
	Delete(ctx context.Context, db *gorm.DB, id uint64) error
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) New(cxt context.Context, db *gorm.DB, user *models.User) error {
	return db.WithContext(cxt).Create(user).Error
}

func (r *userRepositoryImpl) Get(cxt context.Context, id uint64) *models.User {
	var user models.User
	tx := db.WithContext(cxt).First(&user, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}
	return &user
}

func (r *userRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uint64) error {
	return db.WithContext(ctx).Delete(&models.User{ID: id}).Error
}
