package handlers

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/models"
	"github.com/weitien/admin/response"
)

var m sync.Once

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	var instance *UserHandler
	m.Do(func() {
		instance = &UserHandler{}
	})
	return instance
}

func (u *UserHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/user", u.GetUser)
	r.POST("/user", u.Create)
}

func (u *UserHandler) Create(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.Error(err)
		return
	}
}

func (u *UserHandler) GetUser(c *gin.Context) {
	c.Set(response.DATA_KEY, "Weitien")
}
