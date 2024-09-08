package handlers

import (
	"errors"
	"github.com/walterwong1001/admin/internal/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/walterwong1001/admin/internal/machine"
	"github.com/walterwong1001/gin_common_libs/pkg/page"
	"github.com/walterwong1001/gin_common_libs/pkg/response"
)

type AbstractHandler struct {
}

func getPathParamAsInt(c *gin.Context, key string) (uint64, error) {
	return strconv.ParseUint(c.Param(key), 10, 64)
}

func abort(c *gin.Context, err error) {
	abortWithMessage(c, err, "")
}

func abortWithMessage(c *gin.Context, err error, text string) {
	if text != "" {
		err = errors.New(text)
	}
	_ = c.Error(err)
	c.Abort()
}

func nextId() uint64 {
	return machine.NextID()
}

func createTime() int64 {
	return time.Now().UnixMilli()
}

func render(c *gin.Context, data any) {
	c.Set(response.DATA_KEY, data)
}

// paginator 分页器
func paginator[T any](c *gin.Context) (page.Paginator[T], error) {
	current, err := getPathParamAsInt(c, "current")
	if err != nil {
		abort(c, err)
		return nil, err
	}
	size, err := getPathParamAsInt(c, "size")
	if err != nil {
		abort(c, err)
		return nil, err
	}
	return page.NewPagination[T](int(current), int(size)), nil
}

// queryParams 查询参数
func queryParams[T any](c *gin.Context) (T, error) {
	var filter T
	if err := c.ShouldBindQuery(&filter); err != nil {
		abort(c, err)
		return filter, err
	}
	return filter, nil
}

// pagination 抽象分页逻辑
func pagination[T, F any](c *gin.Context, service services.Paginator[T, F]) {
	// 分页信息
	p, err := paginator[T](c)
	if err != nil {
		abort(c, err)
		return
	}
	// 过滤参数
	filter, err := queryParams[F](c)
	if err != nil {
		abort(c, err)
		return
	}

	err = service.Pagination(c.Request.Context(), p, filter)
	if err != nil {
		abort(c, err)
		return
	}
	render(c, p)
}
