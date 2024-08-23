package middleware

import (
	"context"
	"github.com/weitien/admin/global"
	"github.com/weitien/admin/pkg/response"
	"github.com/weitien/admin/pkg/token"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/weitien/admin/internal/services"
	"github.com/weitien/admin/pkg/trie"
)

var cache = trie.New[meta]()

type meta struct {
	allowed bool
	roles   []uint64
}

func Authorization() gin.HandlerFunc {
	initAccessCache()
	return func(c *gin.Context) {
		//获取 Authorization 头部信息，解析token
		claims, err := token.ParseJWT(c.GetHeader("Authorization"), global.CONFIG.Jwt.SecretKey)
		if err != nil {
			log.Printf("%v", err)
		}
		// 拆分Path
		segments := append(strings.Split(strings.Replace(c.Request.URL.Path, "/api/", "", 1), "/"), c.Request.Method)
		// 匹配权限
		_meta := cache.Find(segments)

		if _meta != nil {
			// 白名单
			if _meta.allowed {
				c.Next()
				return
			}
			if claims != nil && claims.Roles != nil {
				// 对比token中的role和权限中绑定的role
				for _, r1 := range _meta.roles {
					for _, r2 := range claims.Roles {
						if r1 == r2 {
							c.Next()
							return
						}
					}
				}
			}
			// 将用户ID保存在上下文中
			c.Set("CURRENT_USER_ID", claims.ID)
		}

		c.JSON(http.StatusForbidden, response.Error(http.StatusForbidden, "Forbidden"))
		c.Abort()
		return
	}
}

func initAccessCache() {
	ctx := context.Background()
	m := make(map[uint64][]uint64)
	// 获取角色权限对应关系, 并根据权限ID分组
	for _, item := range services.NewRolePermissionService().All(ctx) {
		if _, exists := m[item.PermissionId]; !exists {
			m[item.PermissionId] = []uint64{item.RoleId}
		} else {
			m[item.PermissionId] = append(m[item.PermissionId], item.RoleId)
		}
	}

	// 获取所以权限信息，并添加到前缀树
	for _, p := range services.NewPermissionService().All(ctx) {
		segments := append(strings.Split(p.Path, "/"), p.Method)
		if strings.HasPrefix(p.Path, "/") {
			segments = segments[1:]
		}
		_meta := meta{
			allowed: p.IsAllowed(),
		}
		if !p.IsAllowed() {
			_meta.roles = m[p.ID]
		}
		cache.Add(segments, _meta)
	}
}
