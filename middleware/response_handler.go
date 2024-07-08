package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/response"
)

// GlobalResponseHandler 全局统一响应Handler
func GlobalResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.IsAborted() {
			return
		}

		// 发生异常，获取最后一个异常
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, err.Error()))
			return
		}

		data, _ := c.Get(response.DATA_KEY)
		c.JSON(http.StatusOK, response.Success(data))
	}
}

// NoRoute 404 错误处理器
func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, response.R{
		Code:    http.StatusNotFound,
		Message: "Not found",
	})

	c.Abort()
}

// NoMethod 方法不匹配处理器
func NoMethod(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, response.R{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method not allowed",
	})
	c.Abort()
}
