package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/internal/models"
	"github.com/weitien/admin/internal/services"
	"github.com/weitien/admin/pkg/crypto"
)

// 用户名密码认证
type usernamePasswordAuthentication struct {
	service services.AccountServicer
}

func newUsernamePasswordAuthentication() AuthStrategy {
	return &usernamePasswordAuthentication{services.AccountService()}
}

func (s *usernamePasswordAuthentication) Authenticate(c *gin.Context, a *models.AuthCredential) (*models.Account, error) {
	acc := s.service.GetAccountByType(c.Request.Context(), a.Identifier, a.Type)
	if acc == nil {
		return nil, errors.New("invalid user name")
	}
	b := crypto.Matches(acc.Password, a.Secret)
	if b {
		return acc, nil
	}
	return nil, errors.New("user name and password do not match")
}
