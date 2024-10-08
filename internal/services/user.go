package services

import (
	"context"

	"github.com/walterwong1001/admin/internal/machine"
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/repositories"
	"gorm.io/gorm"
)

type UserService interface {
	New(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id uint64) *models.User
	Delete(ctx context.Context, id uint64) error
	All(ctx context.Context) []*models.User
	UserInfo(ctx context.Context, id uint64) *models.UserInfo
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
	return s.repository.Get(ctx, repositories.GetDB(), id)
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

func (s *userServiceImpl) All(ctx context.Context) []*models.User {
	return s.repository.All(ctx, repositories.GetDB())
}

func (s *userServiceImpl) UserInfo(ctx context.Context, id uint64) *models.UserInfo {
	u := s.repository.Get(ctx, repositories.GetDB(), id)
	if u == nil {
		return nil
	}
	return &models.UserInfo{
		ID:     id,
		Name:   u.Name,
		Email:  u.Email,
		Mobile: u.Mobile,
	}
}
