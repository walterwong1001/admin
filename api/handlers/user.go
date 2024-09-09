package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/global"
	"github.com/walterwong1001/admin/internal/models"
	"github.com/walterwong1001/admin/internal/services"
	"github.com/walterwong1001/gin_common_libs/crypto"
	. "github.com/walterwong1001/gin_common_libs/endpoints"
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
	r.GET("/user/all", h.All)
	r.GET("/user/current", h.CurrentUserInfo)
}

func (h *userHandler) New(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		_ = c.Error(err)
		return
	}
	ciphertext, err := crypto.Encode(user.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}
	user.ID = NextId()
	user.Password = ciphertext
	user.CreateTime = CreateTime()

	if err = h.service.New(c.Request.Context(), &user); err != nil {
		_ = c.Error(err)
		c.Abort()
	}
}

func (h *userHandler) Get(c *gin.Context) {
	id, err := PathParamAsInt(c, "id")
	if err != nil {
		Abort(c, err)
		return
	}
	Render(c, h.service.Get(c.Request.Context(), id))
}

func (h *userHandler) Delete(c *gin.Context) {
	id, err := PathParamAsInt(c, "id")
	if err != nil {
		Abort(c, err)
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		AbortWithMessage(c, err, "failed to delete user")
	}
}

func (h *userHandler) All(c *gin.Context) {
	Render(c, h.service.All(c.Request.Context()))
}

func (h *userHandler) CurrentUserInfo(c *gin.Context) {
	v, exists := c.Get(global.KEY_CURRENT_USER_ID)
	if !exists {
		Abort(c, errors.New("user not sign in"))
		return
	}
	s, ok := v.(string)
	if !ok {
		Abort(c, errors.New("invalid user"))
		return
	}
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		Abort(c, errors.New("invalid user id"))
		return
	}
	info := h.service.UserInfo(c.Request.Context(), uint64(id))

	Render(c, info)
}
