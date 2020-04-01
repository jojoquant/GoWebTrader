package middleware

import (
	"GoWebTrader/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "请求头中auth为空",
			})
			c.Abort()
			return
		}

		mc, err := handler.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				// "code":2005,
				"msg": "无效的Token",
			})
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c中
		c.Set("username", mc.Username)
		c.Next()
	}
}
