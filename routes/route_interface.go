package routes

import "github.com/gin-gonic/gin"

type IRoute interface {
	RegisterRoutes(route *gin.RouterGroup)
}
