package logic

import (
	"webapp/dao/mysql"
	"webapp/global"
	"webapp/models"
)

func SignUp(ps *models.ParamSignUp) {
	// 判断用户名是否可用
	mysql.QueryUserByUsername()
	// 生成UID
	global.Snflk.GetID()
	// 用户密码加密
	// 保存进数据库
	mysql.InsertUser()
}
