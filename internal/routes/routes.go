package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/api/handlers"
	"github.com/walterwong1001/gin_common_libs/router"
)

func RegisterRoutes(engine *gin.Engine) {

	group := engine.Group("/api")
	routers := []router.Router{
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
