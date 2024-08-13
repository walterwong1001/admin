package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/pkg/models"
	"github.com/weitien/admin/pkg/services"
	"github.com/weitien/admin/plugin/trie"
	"log"
	"strings"
)

var permissions []*models.Permission

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头部信息
		//token := c.GetHeader("Authorization")
		url := c.Request.URL
		log.Println(url)
		c.Next()
	}
}

func allPermissions() {
	service := services.NewPermissionService()
	arr := service.All(context.Background())
	t := trie.New[[]uint64]()
	for _, p := range arr {
		s := strings.Split(p.Path, "/")
		if strings.HasPrefix(p.Path, "/") {
			s = s[1:]
		}
		t.Add(s, []uint64{123456})
	}
}
