package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/services"
	"github.com/walterwong1001/admin/page"
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
	r.GET("/role/:current/:pageSize", h.Pagination)
}

func (h *roleHandler) New(c *gin.Context) {
	var obj models.Role
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

func (h *roleHandler) Delete(c *gin.Context) {
	id, err := getPathParamAsInt(c, "id")
	if err != nil {
		abort(c, err)
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		abortWithMessage(c, err, "failed to delete role")
	}
}

func (h *roleHandler) Update(c *gin.Context) {
	var obj models.Role
	if err := c.Bind(&obj); err != nil {
		abort(c, err)
		return
	}

	if obj.ID == 0 {
		abort(c, errors.New("ID not specified"))
		return
	}

	if err := h.service.Update(c.Request.Context(), &obj); err != nil {
		abortWithMessage(c, err, "failed to update role")
	}
}

func (h *roleHandler) All(c *gin.Context) {
	render(c, h.service.All(c.Request.Context()))
}

func (h *roleHandler) Pagination(c *gin.Context) {
	current, err := getPathParamAsInt(c, "current")
	if err != nil {
		abort(c, err)
		return
	}
	pageSize, err := getPathParamAsInt(c, "pageSize")
	if err != nil {
		abort(c, err)
		return
	}

	p := page.NewPagination[*models.Role](int(current), int(pageSize))

	var filter models.RoleFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		abort(c, err)
		return
	}

	err = h.service.Pagination(c.Request.Context(), p, &filter)
	if err != nil {
		abort(c, err)
		return
	}
	render(c, p)
}
