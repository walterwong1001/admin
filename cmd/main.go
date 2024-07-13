package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/weitien/admin/snowflake"
	_ "github.com/weitien/admin/utils"

	"github.com/weitien/admin/global"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/middleware"
	"github.com/weitien/admin/response"
	"github.com/weitien/admin/routes"
)

var machineId uint16

func main() {
	r := gin.Default()

	machineId = snowflake.InitSnowFlake(global.CONFIG.Snowflake.Register, global.CONFIG.Application)

	if err := response.InitValidatorTranslator(global.CONFIG.Locale); err != nil {
		log.Println("Init validator translator failed")
	}

	r.Use(middleware.RequestElapsedHandler(), middleware.GlobalResponseHandler())
	r.HandleMethodNotAllowed = true
	r.NoRoute(middleware.NoRoute)
	r.NoMethod(middleware.NoMethod)

	// 加载路由配置
	routes.RegisterRoutes(r)

	// Create the server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.CONFIG.Server.Port),
		Handler: r,
	}
	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// Print the current process ID
	fmt.Printf("Current process ID: %d\n", os.Getpid())

	gracefulShutdown(srv)
}

func gracefulShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Service shutting down...")

	// Context with timeout to allow graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// Unregister the snowflake machine
	snowflake.GetMachineRegister(global.CONFIG.Snowflake.Register).Unregister(machineId)
	log.Println("Server exiting")
}
