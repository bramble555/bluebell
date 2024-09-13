package controllers

import (
	"bluebell/global"
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CommunityHandler 查询到所有社区ID
func CommunityHandler(c *gin.Context) {
	// 查询到所有社区(community_id,community_name)，以切片形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		global.Log.Error("logic.GetCommunityList() failed", err.Error())
		ResponseError(c, CodeServerBusy) // 不轻易把服务端错误暴露给外面
		return
	}
	ResponseSucceed(c, data)
}

// CommunityDetailByIDHandler 根据ID查询社区详情
func CommunityDetailByIDHandler(c *gin.Context) {
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		global.Log.Error("logic.GetCommunityDetailByID() failed", err.Error())
		ResponseError(c, CodeInvalidID)
		return
	}
	communityList, err := logic.GetCommunityDetailByID(ID)
	if err != nil {
		global.Log.Error("logic.GetCommunityDetailByID() failed", err.Error())
		ResponseError(c, CodeInvalidID)
		return
	}
	ResponseSucceed(c, communityList)
}
