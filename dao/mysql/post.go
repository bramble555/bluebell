package mysql

import (
	"bluebell/global"
	"bluebell/models"
)

func CreatePost(p *models.Post) error {
	sqlStr := ` insert into post(
	post_id,title,content,author_id,community_id
	) values (?,?,?,?,?)
	`
	_, err := global.DB.Exec(sqlStr, p.PostID, p.Title, p.Content, p.UserID, p.CommunityID)
	return err
}

func GetPostDetail(p *models.Post) error {
	sqlStr := ` select post_id, author_id, community_id, title, content, create_time
	from post
	where post_id = ?
	`
	return global.DB.QueryRow(sqlStr, p.PostID).Scan(
		&p.PostID,
		&p.UserID,
		&p.CommunityID,
		&p.Title,
		&p.Content,
		&p.CreateTime,
	)
}
