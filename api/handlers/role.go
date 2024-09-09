package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/services"
	. "github.com/walterwong1001/gin_common_libs/endpoints"
)

type roleHandler struct {
	service services.RoleService
}

func RoleHandler() *roleHandler {
	return &roleHandler{service: services.NewRoleService()}
}

func (h *roleHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/role", h.New)
	r.DELETE("/role/:id", h.Delete)
	r.PUT("/role", h.Update)
	r.GET("/role/all", h.All)
	r.GET("/role/:current/:size", h.Pagination)
}

func (h *roleHandler) New(c *gin.Context) {
	var obj models.Role
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

func (h *roleHandler) Delete(c *gin.Context) {
	id, err := PathParamAsInt(c, "id")
	if err != nil {
		Abort(c, err)
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		AbortWithMessage(c, err, "failed to delete role")
	}
}

func (h *roleHandler) Update(c *gin.Context) {
	var obj models.Role
	if err := c.Bind(&obj); err != nil {
		Abort(c, err)
		return
	}

	if obj.ID == 0 {
		Abort(c, errors.New("ID not specified"))
		return
	}

	if err := h.service.Update(c.Request.Context(), &obj); err != nil {
		AbortWithMessage(c, err, "failed to update role")
	}
}

func (h *roleHandler) All(c *gin.Context) {
	Render(c, h.service.All(c.Request.Context()))
}

func (h *roleHandler) Pagination(c *gin.Context) {
	Pagination[*models.Role, *models.RoleFilter](c, h.service, nil)
}
