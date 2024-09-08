package logic

import (
	"errors"
	"webapp/dao/mysql"
	"webapp/global"
	"webapp/models"
)

func SignUp(ps *models.ParamSignUp) error {
	// 判断用户名是否可用
	ok, err := mysql.CheckUserExist(ps.Username)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("用户已存在,请换个用户名")
	}
	// 生成UID
	userID := global.Snflk.GetID()
	// 构造一个User实例，插入数据库
	u := models.User{
		UserID:   userID,
		Username: ps.Username,
		Password: ps.Password,
	}
	// 用户密码加密
	// 保存进数据库
	err = mysql.InsertUser(&u)
	if err != nil {
		return err
	}
	return nil
}
