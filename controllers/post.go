package controllers

import (
	"bluebell/global"
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	// 获取参数以及参数校验   (用户数据给到前端，前端交给服务端，然后创建数据库)
	var p = new(models.Post)
	err := c.ShouldBind(p)
	if err != nil {
		global.Log.Errorln("ShouldBind err", err.Error())
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 创建帖子

	// 获取p的作者ID,也就是当前用户ID
	p.UserID, err = getCurUserID(c)
	if err != nil {
		global.Log.Errorln("logic CreatePost err", err.Error())
		ResponseError(c, CodeNeedLogin)
		return
	}
	err = logic.CreatePost(p)
	if err != nil {
		global.Log.Errorln("logic CreatePost err", err.Error())
		ResponseError(c, CodeInvalidParam)
		return
	}

	//返回响应
	ResponseSucceed(c, CodeSucceed)
}

func GetPostDetailHandler(c *gin.Context) {
	// 获取参数(从URL中获取帖子的ID)
	p := new(models.Post)
	pid := c.Param("id")
	id, err := strconv.ParseInt(pid, 10, 64)
	p.PostID = int(id)
	if err != nil {
		global.Log.Errorln("Params error", err.Error())
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据id取出帖子数据
	pd, err := logic.GetPostDetail(p)
	if err != nil {
		global.Log.Errorln("logic.GetPostDetail error", err.Error())
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 返回响应
	ResponseSucceed(c, pd)
}
