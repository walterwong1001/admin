package services

import (
	"context"

	"github.com/weitien/admin/repositories"
)

type AccountServicer interface {
	ChangeAccountStatus(ctx context.Context, id uint64, status uint)
}

type accountService struct {
	repository repositories.AccountRepository
}

func NewAccountService() AccountServicer {
	return &accountService{repositories.NewAccountRepository()}
}

func (s *accountService) ChangeAccountStatus(ctx context.Context, id uint64, status uint) {
	s.repository.ChangeAccountStatus(ctx, repositories.GetDB(), id, status)
}
