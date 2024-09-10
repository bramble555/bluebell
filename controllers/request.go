package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurUserID获取当前用户ID
func getCurUserID(c *gin.Context) (userID int, err error) {
	uid, ok := c.Get("userID")
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
