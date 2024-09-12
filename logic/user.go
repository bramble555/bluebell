package logic

import (
	"errors"
	"bluebell/dao/mysql"
	"bluebell/global"
	"bluebell/models"
	"bluebell/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
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
	u.Password, err = hashPassword(u.Password)
	// 生成的密码是59~72位，当然也不能接收超越72位的密码，所以需要数据库varchar(72)
	if err != nil {
		return err
	}

	// 保存进数据库
	err = mysql.InsertUser(&u)
	if err != nil {
		return err
	}
	return nil
}
func Login(pl *models.ParamLogin) (string, error) {
	// 判断用户名是否存在
	u := &models.User{
		Username: pl.Username,
		Password: pl.Password,
	}
	ok, err := mysql.CheckUserExist(u.Username)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("用户名不存在")
	}
	// 判断密码是否错误
	hashedPassword, err := mysql.QueryPassword(u)
	if err != nil {
		return "", err
	}
	err = comparePasswords(hashedPassword, u.Password)
	if err != nil {
		// 密码不一致的错误
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return "", errors.New("密码错误")
		}
		// 其他错误
		return "", err
	}
	// 用户名和密码都正确，生成token
	return jwt.GenToken(u.UserID)
}
func hashPassword(password string) (string, error) {
	// 使用bcrypt库的GenerateFromPassword函数进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func comparePasswords(hashedPassword, inputPassword string) error {
	// 使用bcrypt库的CompareHashAndPassword函数比较密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err
}
