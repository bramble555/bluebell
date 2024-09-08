package controllers

import (
	"webapp/global"
	"webapp/logic"
	"webapp/models"

	"github.com/gin-gonic/gin"
)

func SignUphandler(c *gin.Context) {
	// 获取参数并且参数校验
	ps := new(models.ParamSignUp)
	if err := c.ShouldBind(ps); err != nil {
		global.Log.Infoln(ps)
		global.Log.Error("SignUp with invaild params,they are not json parms", err.Error())
		c.JSON(200, gin.H{
			"msg": "请求参数有误",
		})
		return
	}

	// 业务处理
	logic.SignUp(ps)
	// 返回响应
	c.String(200, "ok")
}
