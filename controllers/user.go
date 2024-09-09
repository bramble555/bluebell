package controllers

import (
	"webapp/global"
	"webapp/logic"
	"webapp/models"

	"github.com/gin-gonic/gin"
)

// SignUphandler 实现注册功能
func SignUphandler(c *gin.Context) {
	// 获取参数并且参数校验
	ps := new(models.ParamSignUp)
	if err := c.ShouldBind(ps); err != nil {
		global.Log.Error("SignUp with invaild params,they are not json params", err.Error())
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 业务处理
	err := logic.SignUp(ps)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 返回响应
	c.JSON(200, gin.H{
		"msg": "succeed signup",
	})
}

// Loginhandler 实现登录功能
func Loginhandler(c *gin.Context) {
	// 获取参数并且参数校验
	pl := new(models.ParamLogin)
	err := c.ShouldBind(pl)
	if err != nil {
		global.Log.Error("Login with invaild params,they are not json params", err.Error())
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 业务处理
	err = logic.Login(pl)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 返回响应
	c.JSON(200, gin.H{
		"msg": "succeed login",
	})
}
