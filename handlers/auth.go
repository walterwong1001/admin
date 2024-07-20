package handlers

import (
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
	if err := c.Bind(&credential); err != nil {
		c.Error(err)
		return
	}

	strategy := auth.GetAuthStrategy(credential.Type)
	acc, err := strategy.Authenticate(c, &credential)
	if err != nil {
		c.Error(err)
		c.Abort()
	}
	c.Set(response.DATA_KEY, acc)
}
