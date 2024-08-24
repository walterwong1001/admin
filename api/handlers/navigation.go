package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/services"
)

type navigationHandler struct {
	service services.NavigationServicer
}

func NavigationHandler() *navigationHandler {
	return &navigationHandler{service: services.NewNavigationService()}
}

func (h *navigationHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/navigation", h.New)
	r.DELETE("/navigation/:id", h.Delete)
	r.PUT("/navigation", h.Update)
}

func (h *navigationHandler) New(c *gin.Context) {
	var nav models.Navigation
	if err := c.Bind(&nav); err != nil {
		abort(c, err)
		return
	}

	nav.ID = nextId()

	if err := h.service.New(c.Request.Context(), &nav); err != nil {
		abort(c, err)
	}
}

func (h *navigationHandler) Delete(c *gin.Context) {
	id, err := getPathParamAsInt(c, "id")
	if err != nil {
		abort(c, err)
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		abortWithMessage(c, err, "failed to delete navigation")
	}
}

func (h *navigationHandler) Update(c *gin.Context) {
	var nav models.Navigation
	if err := c.Bind(&nav); err != nil {
		abort(c, err)
		return
	}

	if nav.ID == 0 {
		abort(c, errors.New("ID not specified"))
		return
	}

	if err := h.service.Update(c.Request.Context(), &nav); err != nil {
		abortWithMessage(c, err, "failed to update navigation")
	}
}
