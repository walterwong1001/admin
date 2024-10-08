package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/api/handlers"
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
		handlers.UserRoleHandler(),
	}

	for _, r := range routers {
		r.RegisterRoutes(group)
	}
}
