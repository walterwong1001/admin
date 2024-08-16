package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	accesslog "github.com/weitien/admin/plugin/access_log"
)

func AccessLog(logger accesslog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		var requestBody []byte
		if ctx.Request.Body != nil {
			// 读取请求体
			requestBody, _ = io.ReadAll(ctx.Request.Body)
			// 重新设置请求体，以便后续处理
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		ctx.Next()

		// 计算响应时间
		latency := time.Since(startTime).Seconds()
		userId, _ := ctx.Get("CURRENT_USER_ID")
		// 生成日志记录
		metrics := map[string]any{
			"timestamp":            time.Now().UnixMilli(),
			"remote_addr":          ctx.ClientIP(),
			"remote_user":          userId,
			"request_method":       ctx.Request.Method,
			"request_uri":          ctx.Request.RequestURI,
			"server_protocol":      ctx.Request.Proto,
			"status":               ctx.Writer.Status(),
			"body_bytes_sent":      ctx.Writer.Size(),
			"http_referer":         ctx.Request.Referer(),
			"http_user_agent":      ctx.Request.UserAgent(),
			"http_x_forwarded_for": ctx.GetHeader("X-Forwarded-For"),
			"request_time":         latency,
			"request_body":         string(requestBody),
		}
		logger.Log(ctx, metrics)
	}
}
