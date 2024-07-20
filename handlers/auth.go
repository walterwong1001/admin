package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/auth"
	"github.com/weitien/admin/models"
	"github.com/weitien/admin/response"
)

type authHandler struct{}

func AuthHandler() *authHandler {
	return &authHandler{}
}

func (h *authHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/auth", h.Authenticate)
}

func (h *authHandler) Authenticate(c *gin.Context) {
	var credential models.AuthCredential
	if err := c.BindJSON(&credential); err != nil {
		_ = c.Error(err)
		return
	}

	strategy := auth.GetAuthStrategy(credential.Type)
	if strategy == nil {
		_ = c.Error(errors.New("invalid authentication type"))
		c.Abort()
		return
	}
	acc, err := strategy.Authenticate(c, &credential)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.Set(response.DATA_KEY, acc)
}
