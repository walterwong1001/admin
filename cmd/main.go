package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/config"
	"github.com/weitien/admin/middleware"
	"github.com/weitien/admin/routes"
)

func main() {
	r := gin.Default()

	// 初始化配置
	config.Init()

	r.Use(middleware.RequestElapsedHandler(), middleware.GlobalResponseHandler())
	r.HandleMethodNotAllowed = true
	r.NoRoute(middleware.NoRoute)
	r.NoMethod(middleware.NoMethod)

	// 加载路由配置
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
