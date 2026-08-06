package middleware

import (
	"go-server/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "请求头中缺少 Authorization",
			})
			c.Abort() // 重要：终止后续执行
			return
		}

		// 支持 "Bearer token" 格式
		tokenStr := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// 2. 解析和验证 Token
		claims, err := util.ParseToken(tokenStr)
		if err != nil {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "登录已过期或无效，请重新登录",
			})
			c.Abort()
			return
		}

		// 3. 将用户信息存入 Context，供后续 handler 使用
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("claims", claims) // 如果需要更多信息

		c.Next() // 继续执行后续的 handler
	}
}
