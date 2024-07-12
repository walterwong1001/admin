package main

import (
	"fmt"
	"log"

	"github.com/weitien/admin/global"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/middleware"
	_ "github.com/weitien/admin/repositories"
	"github.com/weitien/admin/response"
	"github.com/weitien/admin/routes"
)

func main() {
	r := gin.Default()

	if err := response.InitValidatorTranslator(global.CONFIG.Locale); err != nil {
		log.Println("Init validator translator failed")
	}

	r.Use(middleware.RequestElapsedHandler(), middleware.GlobalResponseHandler())
	r.HandleMethodNotAllowed = true
	r.NoRoute(middleware.NoRoute)
	r.NoMethod(middleware.NoMethod)

	// 加载路由配置
	routes.RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%d", global.CONFIG.Server.Port))
}
