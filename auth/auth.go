package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/models"
)

var strategies = make(map[models.AccountType]AuthStrategy)

func init() {
	fmt.Println("***************************************************")
	strategies[models.AccountTypeUsername] = UsernamePasswordAuthentication()
}

type AuthStrategy interface {
	Authenticate(c *gin.Context, p *models.AuthCredential) (*models.Account, error)
}

func GetAuthStrategy(key models.AccountType) AuthStrategy {
	return strategies[key]
}
