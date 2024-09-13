package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityList 得到社区列表
func GetCommunityList() (*[]models.Community, error) {
	// 查到数据库，得到社区列表
	return mysql.GetCommunityList()
}

// GetCommunityDetailByID 得到某个社区详情
func GetCommunityDetailByID(id uint64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
