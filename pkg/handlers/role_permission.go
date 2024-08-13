package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/pkg/models"
	"github.com/weitien/admin/pkg/services"
)

type rolePermissionHandler struct {
	service services.RolePermissionService
}

func RolePermissionHandler() *rolePermissionHandler {
	return &rolePermissionHandler{service: services.NewRolePermissionService()}
}

func (h *rolePermissionHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/role_permission/binding", h.Binding)
	r.DELETE("/role_permission/binding", h.Binding)
}

func (h *rolePermissionHandler) Binding(c *gin.Context) {
	var obj models.RolePermission

	if err := c.Bind(&obj); err != nil {
		abort(c, err)
		return
	}

	method := c.Request.Method

	if method == http.MethodPost {
		if err := h.service.Bind(c.Request.Context(), &obj); err != nil {
			abort(c, err)
		}
	}

	if method == http.MethodDelete {
		if err := h.service.UnBind(c.Request.Context(), &obj); err != nil {
			abort(c, err)
		}
	}
}
