package routers

import (
	"webapp/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRounter(mode string) *gin.Engine {
	// 如果是发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(mode)
	}
	r := gin.Default()
	// 注册业务路由
	r.POST("/signup", controllers.SignUphandler)
	// 登录业务路由
	r.POST("/login", controllers.Loginhandler)
	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}
