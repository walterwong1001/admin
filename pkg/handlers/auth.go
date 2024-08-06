package handlers

import (
	"errors"
	"github.com/weitien/admin/pkg/global"
	"github.com/weitien/admin/pkg/models"
	"github.com/weitien/admin/pkg/response"
	"github.com/weitien/admin/pkg/token"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/auth"
)

type authHandler struct{}

var jwt = global.CONFIG.Jwt

func AuthHandler() *authHandler {
	return &authHandler{}
}

func (h *authHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/auth", h.Authenticate)
}

func (h *authHandler) Authenticate(c *gin.Context) {
	var credential models.AuthCredential
	if err := c.BindJSON(&credential); err != nil {
		abort(c, err)
		return
	}

	strategy := auth.GetAuthStrategy(credential.Type)
	if strategy == nil {
		abort(c, errors.New("invalid authentication type"))
		return
	}
	acc, err := strategy.Authenticate(c, &credential)
	if err != nil {
		abort(c, err)
		return
	}
	sub := strconv.Itoa(int(acc.UserID))
	jwt, err := token.NewJWT(sub, sub, jwt.Days, jwt.SecretKey, jwt.Issuer)
	if err != nil {
		abort(c, err)
		return
	}
	c.Set(response.DATA_KEY, jwt)
}
