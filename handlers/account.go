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

func NewAccountHandler() *accountHandler {
	return &accountHandler{services.NewAccountService()}
}

func (h *accountHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.PUT("/account/lock/:id", h.LockAccount)
	r.PUT("/account/activity/:id", h.ActiviteAccount)
}

func (h *accountHandler) LockAccount(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err)
		c.Abort()
	}
	h.service.ChangeAccountStatus(c.Request.Context(), id, models.AccountLocked)
}

func (h *accountHandler) ActiviteAccount(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err)
		c.Abort()
	}
	h.service.ChangeAccountStatus(c.Request.Context(), id, models.AccountActivity)
}
