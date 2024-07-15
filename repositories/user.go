package repositories

import (
	"context"
	"gorm.io/gorm"

	"github.com/weitien/admin/models"
)

type UserRepository interface {
	CreateUser(cxt context.Context, db *gorm.DB, user *models.User) error
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) CreateUser(cxt context.Context, db *gorm.DB, user *models.User) error {
	return db.WithContext(cxt).Create(user).Error
}
