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
