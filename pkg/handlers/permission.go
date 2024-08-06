package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/pkg/models"
	"github.com/weitien/admin/pkg/services"
)

type permissionHandler struct {
	service services.PermissionService
}

func PermissionHandler() *permissionHandler {
	return &permissionHandler{service: services.NewPermissionService()}
}

func (h *permissionHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/permission", h.New)
	r.DELETE("/permission/:id", h.Delete)
	r.PUT("/permission", h.Update)
}

func (h *permissionHandler) New(c *gin.Context) {
	var obj models.Permission
	if err := c.Bind(&obj); err != nil {
		abort(c, err)
		return
	}

	obj.ID = nextId()
	obj.CreateTime = createTime()

	if err := h.service.New(c.Request.Context(), &obj); err != nil {
		abort(c, err)
	}
}

func (h *permissionHandler) Delete(c *gin.Context) {
	id, err := getPathParamAsInt(c, "id")
	if err != nil {
		abort(c, err)
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		abortWithMessage(c, err, "failed to delete permission")
	}
}

func (h *permissionHandler) Update(c *gin.Context) {
	var obj models.Permission
	if err := c.Bind(&obj); err != nil {
		abort(c, err)
		return
	}

	if obj.ID == 0 {
		abort(c, errors.New("ID not specified"))
		return
	}

	if err := h.service.Update(c.Request.Context(), &obj); err != nil {
		abortWithMessage(c, err, "failed to update permission")
	}
}
