package controllers

import (
	"bluebell/global"
	"bluebell/logic"

	"github.com/gin-gonic/gin"
)

// CommunityHandler 实现社区功能
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
