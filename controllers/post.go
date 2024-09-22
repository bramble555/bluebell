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

// GetPostListHandler 实现Post列表查询
func GetPostListHandler(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var page int
	page64, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		global.Log.Errorln("controller page error", err.Error())
		page64 = 0
	}
	page = int(page64)
	var size int
	size64, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		global.Log.Errorln("controller size error", err.Error())
		page64 = 2
	}
	size = int(size64)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		global.Log.Errorln("logic GetPostList error", err.Error())
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSucceed(c, data)
}

// PostVoteHandler 实现帖子投票功能
func PostVoteHandler(c *gin.Context) {
	// 参数校验
	pv := new(models.ParamPostVote)
	err := c.ShouldBindJSON(pv)
	if err != nil {
		global.Log.Errorln("controller PostVoteHandler err", err.Error())
		ResponseErrorWithData(c, 200, err.Error())
		return
	}
	userID, err := getCurUserID(c)
	if err != nil {
		global.Log.Errorln("controller getCurUserID err", err.Error())
		ResponseErrorWithData(c, 200, err.Error())
		return
	}
	// 逻辑处理
	err = logic.VoteForPost(userID, pv)
	if err != nil {
		global.Log.Errorln("logic VoteForPost err", err.Error())
		ResponseErrorWithData(c, 200, err.Error())
		return
	}
	// 返回响应
	ResponseSucceed(c, pv)

}
