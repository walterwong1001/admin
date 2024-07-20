package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/models"
)

// 策略容器
var strategies = make(map[models.AccountType]AuthStrategy)

// 初始化认证策略
func init() {
	strategies[models.AccountTypeUsername] = newUsernamePasswordAuthentication()
}

// AuthStrategy 认证策略接口
type AuthStrategy interface {
	Authenticate(c *gin.Context, p *models.AuthCredential) (*models.Account, error)
}

// GetAuthStrategy 获取认证策略
func GetAuthStrategy(key models.AccountType) AuthStrategy {
	return strategies[key]
}
