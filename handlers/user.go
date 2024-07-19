package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/machine"
	"github.com/weitien/admin/models"
	"github.com/weitien/admin/response"
	"github.com/weitien/admin/services"
	"github.com/weitien/admin/utils"
)

type userHandler struct {
	Service services.UserService
}

func UserHandler() *userHandler {
	return &userHandler{
		Service: services.NewUserService(),
	}
}

func (h *userHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/user/:id", h.GetUser)
	r.POST("/user", h.CreateUser)
	r.DELETE("/user/:id", h.DeleteUser)
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.Error(err)
		return
	}
	ciphertext, err := utils.Encode(user.Password)
	if err != nil {
		c.Error(err)
		return
	}
	user.ID = machine.GetSnowflake().NextID()
	user.Password = ciphertext
	user.CreateTime = time.Now().UnixMilli()

	err = h.Service.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.Error(err)
		c.Abort()
	}
}

func (h *userHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err)
		c.Abort()
	}
	c.Set(response.DATA_KEY, h.Service.GetUser(c.Request.Context(), id))
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err)
		c.Abort()
	}
	h.Service.DeleteUser(c.Request.Context(), id)
}
