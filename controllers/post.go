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


// GetPostListHandler 实现Post列表查询 升级版
// 根据前端传来的参数动态获取帖子列表
// 按创建时间排序或者分数排序
func GetPostListHandler(c *gin.Context) {
	// 1. 获取参数
	ppl := new(models.ParamPostList)
	defaultPage := 1
	defaultSize := 10
	defaultOrder := redis.OrderByTime // 默认按时间排序
	// 设置默认值
	ppl.Page = defaultPage
	ppl.Size = defaultSize
	ppl.Order = defaultOrder
	// 从URL获取参数
	err := c.ShouldBindQuery(&ppl)
	if err != nil {
		global.Log.Errorf("controller GetPostListHandler2 error: %v", err)
		ResponseErrorWithData(c, CodeInvalidParam, err.Error())
		return
	}

	// 验证参数
	if ppl.Page <= 0 {
		ppl.Page = defaultPage
	}
	if ppl.Size <= 0 {
		ppl.Size = defaultSize
	}
	if ppl.Order == "" {
		ppl.Order = defaultOrder
	}
	// 2. 去redis查询id列表
	// 3. 根据id去数据库查询帖子详情
	// 2 和 3 都在logic去做
	data, err := logic.GetPostList(ppl)
	if err != nil {
		global.Log.Error("logic GetPostList error", err.Error())
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
