package services

import (
	"context"
	"time"

	"github.com/weitien/admin/models"
	"github.com/weitien/admin/repositories"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type userServiceImpl struct {
	repository        repositories.UserRepository
	accountRepository repositories.AccountRepository
}

func NewUserService() UserService {
	return &userServiceImpl{
		repository:        repositories.NewUserRepository(),
		accountRepository: repositories.NewAccountRepository(),
	}
}

func (s *userServiceImpl) CreateUser(ctx context.Context, user *models.User) (err error) {
	tx := repositories.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	// 创建用户
	if err = s.repository.CreateUser(ctx, user); err != nil {
		return err
	}

	// 创建默认账户
	if err = s.accountRepository.NewAccounts(ctx, s.getDefaultAccounts(user)); err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) getDefaultAccounts(u *models.User) []*models.Account {
	// Username Account
	acc := models.Account{
		ID:         1,
		UserID:     u.ID,
		Identifier: u.Name,
		Password:   "123",
		Type:       models.AccountTypeUsername,
		Status:     models.AccountActivity,
		CreateTime: time.Now().UnixMilli(),
	}
	// Email Account
	emailAcc := acc
	emailAcc.ID = 2
	emailAcc.Identifier = u.Email
	emailAcc.Type = models.AccountTypeEmail
	// Mobile Account
	mobileAcc := acc
	mobileAcc.ID = 3
	mobileAcc.Identifier = u.Mobile
	mobileAcc.Type = models.AccountTypeMobile

	return []*models.Account{&acc, &emailAcc, &mobileAcc}
}
