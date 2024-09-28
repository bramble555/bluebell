package routers

import (
	"bluebell/controllers"
	"bluebell/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRounter(mode string) *gin.Engine {
	// 如果是发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(mode)
	}
	r := gin.Default()
	v1 := r.Group("/api/v1")
	// 注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	// 登录业务路由
	v1.POST("/login", controllers.LoginHandler)
	// 权限认证
	v1.Use(middlewares.JWTAuthorMiddleware())
	// 实现社区功能
	v1.GET("/community", controllers.CommunityHandler)
	v1.GET("/community/:id", controllers.CommunityDetailByIDHandler)
	v1.POST("/post", controllers.CreatePostHandler)
	v1.GET("/post/:id", controllers.GetPostDetailHandler)
	v1.GET("/posts", controllers.GetPostListHandler)
	v1.POST("/vote", controllers.PostVoteHandler)
	return r
}
