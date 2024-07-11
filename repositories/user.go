package repository

import (
	"context"

	"github.com/weitien/admin/models"
)

type UserRepository interface {
	CreateUser(cxt context.Context, user *models.User) error
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) CreateUser(cxt context.Context, user *models.User) error {
	return DB.WithContext(cxt).Create(user).Error
}
