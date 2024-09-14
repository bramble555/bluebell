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
	_, err := global.DB.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}
