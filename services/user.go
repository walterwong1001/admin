package services

import (
	"context"

	"github.com/weitien/admin/machine"
	"gorm.io/gorm"

	"github.com/weitien/admin/models"
	"github.com/weitien/admin/repositories"
)

type UserService interface {
	New(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id uint64) *models.User
	Delete(ctx context.Context, id uint64) error
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

func (s *userServiceImpl) New(ctx context.Context, user *models.User) error {
	return repositories.GetDB().Transaction(func(tx *gorm.DB) error {
		// 创建用户
		if err := s.repository.New(ctx, tx, user); err != nil {
			return err
		}
		// 创建默认账户
		if err := s.accountRepository.NewAccounts(ctx, tx, s.getDefaultAccounts(user)); err != nil {
			return err
		}
		return nil
	})
}

func (s *userServiceImpl) Get(ctx context.Context, id uint64) *models.User {
	return s.repository.Get(ctx, id)
}

func (s *userServiceImpl) Delete(ctx context.Context, id uint64) error {
	return repositories.GetDB().Transaction(func(tx *gorm.DB) error {
		err := s.repository.Delete(ctx, tx, id)
		if err != nil {
			return err
		}
		return s.accountRepository.Delete(ctx, tx, id)
	})
}

func (s *userServiceImpl) getDefaultAccounts(u *models.User) []*models.Account {

	snowflake := machine.GetSnowflake()
	// Username Account
	acc := models.Account{
		ID:         snowflake.NextID(),
		UserID:     u.ID,
		Identifier: u.Name,
		Password:   u.Password,
		Type:       models.AccountTypeUsername,
		Status:     models.AccountActivity,
		CreateTime: u.CreateTime,
	}
	// Email Account
	emailAcc := acc
	emailAcc.ID = snowflake.NextID()
	emailAcc.Identifier = u.Email
	emailAcc.Type = models.AccountTypeEmail
	// Mobile Account
	mobileAcc := acc
	mobileAcc.ID = machine.GetSnowflake().NextID()
	mobileAcc.Identifier = u.Mobile
	mobileAcc.Type = models.AccountTypeMobile

	return []*models.Account{&acc, &emailAcc, &mobileAcc}
}
