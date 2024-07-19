package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api")

	handlers.UserHandler().RegisterRoutes(group)

	handlers.SignInHandler().RegisterRoutes(group)

	handlers.NewAccountHandler().RegisterRoutes(group)
}
