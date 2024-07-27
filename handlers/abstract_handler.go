package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/machine"
	"log"
	"strconv"
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
	log.Println(err)
	if text != "" {
		err = errors.New(text)
	}
	_ = c.Error(err)
	c.Abort()
}

func nextId() uint64 {
	return machine.NextID()
}
