package main

import (
	"GoWebTrader/handler"
	"GoWebTrader/middleware"
	"GoWebTrader/setting"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/login", middleware.UserInfo(setting.WebTrader), handler.Login)

	dashboard := r.Group("/dashboard")
	dashboard.Use(middleware.JWTAuth())

	r.Run(fmt.Sprintf(":%d", setting.WebTrader.Port))
}
