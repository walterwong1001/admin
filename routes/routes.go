package routes

import (
	"github.com/gin-gonic/gin"
	h "github.com/weitien/admin/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	group := router.Group("/api")
	routers := []Router{h.AuthHandler(), h.UserHandler(), h.AccountHandler(), h.NavigationHandler(), h.RoleHandler()}

	for _, r := range routers {
		r.RegisterRoutes(group)
	}
}
