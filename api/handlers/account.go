package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/services"
	. "github.com/walterwong1001/gin_common_libs/endpoints"
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
	id, err := PathParamAsInt(c, "id")
	if err != nil {
		Abort(c, err)
		return
	}

	if err := h.service.ChangeAccountStatus(c.Request.Context(), id, models.AccountLocked); err != nil {
		AbortWithMessage(c, err, "failed to lock account")
	}
}

func (h *accountHandler) ActivateAccount(c *gin.Context) {
	id, err := PathParamAsInt(c, "id")
	if err != nil {
		Abort(c, err)
		return
	}

	if err := h.service.ChangeAccountStatus(c.Request.Context(), id, models.AccountActivity); err != nil {
		AbortWithMessage(c, err, "failed to activate account")
	}
}
