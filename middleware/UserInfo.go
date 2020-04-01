package middleware

import (
	"GoWebTrader/setting"

	"github.com/gin-gonic/gin"
)

func UserInfo(w *setting.WebTraderSetting) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("username", w.Username)
		c.Set("password", w.Password)
		c.Next()
	}
}
