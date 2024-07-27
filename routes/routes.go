package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api")

	handlers.AuthHandler().RegisterRoutes(group)

	handlers.UserHandler().RegisterRoutes(group)

	handlers.AccountHandler().RegisterRoutes(group)

	handlers.NavigationHandler().RegisterRoutes(group)
}
