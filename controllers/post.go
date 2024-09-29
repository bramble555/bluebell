package controllers

import (
	"bluebell/dao/redis"
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
	// 参数判断 communtiy_id (目前只有1-4)
	if p.CommunityID < 1 || p.CommunityID > 4 {
		global.Log.Errorf("logic CreatePost err: %s", "community_id参数错误")
		ResponseErrorWithData(c, CodeInvalidParam, "community_id参数错误")
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

// GetPostListFitHandler 获取帖子
func GetPostListFitHandler(c *gin.Context) {
	// ParamPostList 默认值
	ppl := &models.ParamPostList{
		ID:    0,
		Page:  1,
		Size:  10,
		Order: redis.OrderByTime,
	}
	err := c.ShouldBindQuery(&ppl)
	if err != nil {
		global.Log.Errorf("controller GetPostListFitHandler error: %v", err)
		ResponseErrorWithData(c, CodeInvalidParam, err.Error())
		return
	}
	// 参数校验
	if ppl.Page <= 0 {
		ppl.Page = 1
	}
	if ppl.Size <= 0 {
		ppl.Size = 10
	}
	if ppl.Order == "" {
		ppl.Order = redis.OrderByTime
	}
	data, err := logic.GetPostListFit(ppl)
	if err != nil {
		global.Log.Error("logic GetPostListFit error", err.Error())
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
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
