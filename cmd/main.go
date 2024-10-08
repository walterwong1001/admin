package main

import (
	"context"
	// "errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/walterwong1001/admin/global"
	// mdw "github.com/walterwong1001/admin/internal/middleware"
	"github.com/walterwong1001/gin_common_libs/middleware"
	"github.com/walterwong1001/gin_common_libs/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/internal/machine"
	"github.com/walterwong1001/admin/internal/routes"
	"github.com/walterwong1001/admin/internal/services"
)

var machineId uint16
var conf = global.CONFIG

func main() {
	r := gin.Default()

	machineId = machine.InitSnowFlake(conf.Snowflake.Register, conf.Application)

	validator.InitValidator()

	r.Use(middleware.RequestElapsed(), middleware.AccessLog(services.NewAccessLogService()), middleware.GlobalResponse())

	r.HandleMethodNotAllowed = true
	r.NoRoute(middleware.NoRoute)
	r.NoMethod(middleware.NoMethod)

	// 加载路由配置
	routes.RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%d", conf.Server.Port))

	// Create the server
	// srv := &http.Server{
	// 	Addr:    fmt.Sprintf(":%d", conf.Server.Port),
	// 	Handler: r,
	// }
	// // Start the server in a goroutine
	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// // Print the current process ID
	// log.Printf("Current process ID: %d\n", os.Getpid())

	// gracefulShutdown(srv)
}

func gracefulShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutting down the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Service shutting down...")

	// Context with timeout to allow graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutting down the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// Unregister the machine
	machine.GetMachineRegister(conf.Snowflake.Register).Unregister(machineId)
	log.Println("Server exiting")
}
