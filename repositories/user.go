package repositories

import (
	"context"
	"gorm.io/gorm"

	"github.com/weitien/admin/models"
)

type UserRepository interface {
	CreateUser(cxt context.Context, db *gorm.DB, user *models.User) error
	GetUser(cxt context.Context, id uint64) *models.User
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) CreateUser(cxt context.Context, db *gorm.DB, user *models.User) error {
	return db.WithContext(cxt).Create(user).Error
}

func (r *userRepositoryImpl) GetUser(cxt context.Context, id uint64) *models.User {
	var user models.User
	db.First(&user, id)
	return &user
}
