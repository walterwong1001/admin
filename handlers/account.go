package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/models"
	"github.com/weitien/admin/services"
)

type accountHandler struct {
	service services.AccountServicer
}

func AccountHandler() *accountHandler {
	return &accountHandler{services.AccountService()}
}

func (h *accountHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.PUT("/account/lock/:id", h.LockAccount)
	r.PUT("/account/activity/:id", h.ActivateAccount)
}

func (h *accountHandler) LockAccount(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
	}
	h.service.ChangeAccountStatus(c.Request.Context(), id, models.AccountLocked)
}

func (h *accountHandler) ActivateAccount(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
	}
	h.service.ChangeAccountStatus(c.Request.Context(), id, models.AccountActivity)
}
