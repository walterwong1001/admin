package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/services"
	. "github.com/walterwong1001/gin_common_libs/endpoints"
)

type userRoleHandler struct {
	service services.UserRoleService
}

func UserRoleHandler() *userRoleHandler {
	return &userRoleHandler{service: services.NewUserRoleService()}
}

func (h *userRoleHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/user_role/binding", h.Binding)
	r.DELETE("/user_role/binding", h.Binding)
}

func (h *userRoleHandler) Binding(c *gin.Context) {
	var obj models.UserRole

	if err := c.Bind(&obj); err != nil {
		Abort(c, err)
		return
	}

	var err error
	switch c.Request.Method {
	case http.MethodPost:
		err = h.service.Bind(c.Request.Context(), &obj)
	case http.MethodDelete:
		err = h.service.UnBind(c.Request.Context(), &obj)
	default:
		err = errors.New("method not allowed")
	}

	if err != nil {
		Abort(c, err)
	}
}
