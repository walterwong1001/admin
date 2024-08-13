package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/pkg/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api")
	routers := []Router{
		handlers.AuthHandler(),
		handlers.UserHandler(),
		handlers.AccountHandler(),
		handlers.NavigationHandler(),
		handlers.RoleHandler(),
		handlers.PermissionHandler(),
		handlers.RolePermissionHandler(),
	}

	for _, r := range routers {
		r.RegisterRoutes(group)
	}
}
