package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/services"
	. "github.com/walterwong1001/gin_common_libs/endpoints"
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
	r.GET("/permission/all", h.All)
	r.GET("/permission//:current/:size", h.Pagination)
}

func (h *permissionHandler) New(c *gin.Context) {
	var obj models.Permission
	if err := c.Bind(&obj); err != nil {
		Abort(c, err)
		return
	}

	obj.ID = NextId()
	obj.CreateTime = CreateTime()

	if err := h.service.New(c.Request.Context(), &obj); err != nil {
		Abort(c, err)
	}
}

func (h *permissionHandler) Delete(c *gin.Context) {
	id, err := PathParamAsInt(c, "id")
	if err != nil {
		Abort(c, err)
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		AbortWithMessage(c, err, "failed to delete permission")
	}
}

func (h *permissionHandler) Update(c *gin.Context) {
	var obj models.Permission
	if err := c.Bind(&obj); err != nil {
		Abort(c, err)
		return
	}

	if obj.ID == 0 {
		Abort(c, errors.New("ID not specified"))
		return
	}

	if err := h.service.Update(c.Request.Context(), &obj); err != nil {
		AbortWithMessage(c, err, "failed to update permission")
	}
}

func (h *permissionHandler) All(c *gin.Context) {
	arr := h.service.All(c)
	Render(c, arr)
}

func (h *permissionHandler) Pagination(c *gin.Context) {
	Pagination[*models.Permission, *models.PermissionFilter](c, h.service, nil)
}
