package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/pkg/models"
	"github.com/weitien/admin/pkg/services"
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
	id, err := getPathParamAsInt(c, "id")
	if err != nil {
		abort(c, err)
		return
	}

	if err := h.service.ChangeAccountStatus(c.Request.Context(), id, models.AccountLocked); err != nil {
		abortWithMessage(c, err, "failed to lock account")
	}
}

func (h *accountHandler) ActivateAccount(c *gin.Context) {
	id, err := getPathParamAsInt(c, "id")
	if err != nil {
		abort(c, err)
		return
	}

	if err := h.service.ChangeAccountStatus(c.Request.Context(), id, models.AccountActivity); err != nil {
		abortWithMessage(c, err, "failed to activate account")
	}
}
