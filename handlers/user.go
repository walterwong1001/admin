package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/models"
	"github.com/weitien/admin/response"
	"github.com/weitien/admin/services"
	"github.com/weitien/admin/utils"
)

type userHandler struct {
	service services.UserService
}

func UserHandler() *userHandler {
	return &userHandler{
		service: services.NewUserService(),
	}
}

func (h *userHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/user/:id", h.Get)
	r.POST("/user", h.New)
	r.DELETE("/user/:id", h.Delete)
}

func (h *userHandler) New(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		_ = c.Error(err)
		return
	}
	ciphertext, err := utils.Encode(user.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}
	user.ID = nextId()
	user.Password = ciphertext
	user.CreateTime = time.Now().UnixMilli()

	if err = h.service.New(c.Request.Context(), &user); err != nil {
		_ = c.Error(err)
		c.Abort()
	}
}

func (h *userHandler) Get(c *gin.Context) {
	id, err := getPathParamAsInt(c, "id")
	if err != nil {
		abort(c, err)
		return
	}
	c.Set(response.DATA_KEY, h.service.Get(c.Request.Context(), id))
}

func (h *userHandler) Delete(c *gin.Context) {
	id, err := getPathParamAsInt(c, "id")
	if err != nil {
		abort(c, err)
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		abortWithMessage(c, err, "failed to delete user")
	}
}
