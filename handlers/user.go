package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/models"
	"github.com/weitien/admin/response"
	"github.com/weitien/admin/services"
	"github.com/weitien/admin/snowflake"
	"time"
)

type userHandler struct {
	Service services.UserService
}

func UserHandler() *userHandler {
	return &userHandler{
		Service: services.NewUserService(),
	}
}

func (u *userHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/user", u.GetUser)
	r.POST("/user", u.Create)
}

func (u *userHandler) Create(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.Error(err)
		return
	}
	user.ID, _ = snowflake.GetSnowflake().NextID()
	user.CreateTime = time.Now().UnixMilli()
	err := u.Service.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.Error(err)
	}
	c.Abort()
}

func (u *userHandler) GetUser(c *gin.Context) {
	c.Set(response.DATA_KEY, "Weitien")
}
