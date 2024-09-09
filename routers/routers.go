package routers

import (
	"webapp/controllers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	// 注册业务路由
	r.POST("/signup", controllers.SignUphandler)
	// 登录业务路由
	r.POST("/login",controllers.Loginhandler)
	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}
