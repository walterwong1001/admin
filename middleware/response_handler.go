package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/response"
	"log"
	"net/http"
)

// GlobalResponseHandler 全局统一响应Handler
func GlobalResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 发生异常，获取最后一个异常
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			log.Println(err.Error())

			var e *NotRouteOrNoMethodError
			if errors.As(err, &e) {
				// 处理 404 和 405 错误
				c.JSON(http.StatusNotFound, response.Error(e.Code, e.Error()))
				c.Abort()
				return
			}
			// 其他类型错误
			c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, err.Error()))
			return
		}

		data, _ := c.Get(response.DATA_KEY)
		c.JSON(http.StatusOK, response.Success(data))
	}
}

type NotRouteOrNoMethodError struct {
	Code    int
	Message string
}

func (e *NotRouteOrNoMethodError) Error() string {
	return e.Message
}

// NoRoute 404 错误处理器
func NoRoute(c *gin.Context) {
	c.Error(&NotRouteOrNoMethodError{http.StatusNotFound, "Not found"})
	c.Abort()
}

// NoMethod 方法不匹配处理器
func NoMethod(c *gin.Context) {
	c.Error(&NotRouteOrNoMethodError{http.StatusMethodNotAllowed, "Method not allowed"})
	c.Abort()
}
