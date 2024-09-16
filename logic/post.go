package logic

import (
	"bluebell/dao/mysql"
	"bluebell/global"
	"bluebell/models"
)

func CreatePost(p *models.Post) error {
	// 根据雪花算法获取postID
	p.PostID = global.Snflk.GetID()
	return mysql.CreatePost(p)
}

func GetPostDetail(p *models.Post) (*models.PostDetail, error) {
	pd := &models.PostDetail{
		Post: p, // 显式初始化 Post 字段
	}
	err := mysql.GetPostDetail(p)
	if err != nil {
		global.Log.Errorf("logic GetPostDetail err post_id是 %d: %v", p.PostID, err)
		return nil, err
	}
	// 获取社区详情
	cd, err := mysql.GetCommunityDetailByID(p.CommunityID)
	if err != nil {
		global.Log.Errorf("logic GetCommunityDetailByID err community_id是 %d: %v", p.CommunityID, err)
		return nil, err
	}
	pd.CommunityDetail = cd

	// 获取用户详情
	ud, err := mysql.GetUserDetail(p.UserID)
	if err != nil {
		global.Log.Errorf("logic GetUserDetail err user_id是 %d: %v", p.UserID, err)
		return nil, err
	}
	pd.Username = ud.Username

	return pd, nil
}
