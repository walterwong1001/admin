package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api")

	user := handlers.UserHandler()
	user.RegisterRoutes(group)
}
