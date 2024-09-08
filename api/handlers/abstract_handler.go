package handlers

import (
	"errors"
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

func paginator[T any](c *gin.Context) (page.Paginator[T], error) {
	current, err := getPathParamAsInt(c, "current")
	if err != nil {
		abort(c, err)
		return nil, err
	}
	pageSize, err := getPathParamAsInt(c, "size")
	if err != nil {
		abort(c, err)
		return nil, err
	}
	return page.NewPagination[T](int(current), int(pageSize)), nil
}

func queryParams[T any](c *gin.Context) (T, error) {
	var filter T
	if err := c.ShouldBindQuery(&filter); err != nil {
		abort(c, err)
		return filter, err
	}
	return filter, nil
}
