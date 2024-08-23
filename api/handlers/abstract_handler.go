package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/internal/machine"
	"github.com/weitien/admin/pkg/response"
	"strconv"
	"time"
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
