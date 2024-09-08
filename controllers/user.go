package controllers

import (
	"strings"
	"webapp/global"
	"webapp/logic"
	"webapp/models"

	"github.com/gin-gonic/gin"
)

func SignUphandler(c *gin.Context) {
	// 获取参数并且参数校验
	ps := new(models.ParamSignUp)
	if err := c.ShouldBind(ps); err != nil {
		global.Log.Error("SignUp with invaild params,they are not json parms", err.Error())
		c.JSON(200, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	// 上述校验只能校验是否Json，其他逻辑还得自己写
	if strings.Compare(ps.Password, ps.RePassword) != 0 || len(ps.Password) == 0 || len(ps.Username) == 0 {
		global.Log.Error("SingUp with invaild params")
		c.JSON(200, gin.H{
			"msg": "俩次输入密码不一致",
		})
		return
	}
	// 业务处理
	logic.SignUp(ps)
	// 返回响应
	c.String(200, "ok")
}
