package services

import (
	"context"
	"github.com/weitien/admin/pkg/models"
	"github.com/weitien/admin/pkg/repositories"
	"sync"
)

var once sync.Once

type AccountServicer interface {
	ChangeAccountStatus(ctx context.Context, id uint64, status uint8) error
	GetAccountByType(ctx context.Context, identifier string, accType models.AccountType) *models.Account
}

type accountService struct {
	repository repositories.AccountRepository
}

func AccountService() AccountServicer {
	var service AccountServicer
	once.Do(func() {
		service = &accountService{repositories.NewAccountRepository()}
	})
	return service
}

func (s *accountService) ChangeAccountStatus(ctx context.Context, id uint64, status uint8) error {
	return s.repository.ChangeStatus(ctx, repositories.GetDB(), id, status)
}

func (s *accountService) GetAccountByType(ctx context.Context, identifier string, accType models.AccountType) *models.Account {
	return s.repository.GetByType(ctx, repositories.GetDB(), identifier, accType)
}
