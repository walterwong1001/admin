package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestElapsedHandler 请求耗时处理器
func RequestElapsedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 让请求继续执行
		c.Next()

		duration := time.Since(start)

		log.Printf("Request [Path: %s, Method: %s] duration: %v\n", c.Request.URL.Path, c.Request.Method, duration)
	}
}
