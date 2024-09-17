package mysql

import (
	"bluebell/global"
	"bluebell/models"
)

func CreatePost(p *models.Post) error {
	sqlStr := ` insert into post(
	post_id,title,content,user_id,community_id
	) values (?,?,?,?,?)
	`
	_, err := global.DB.Exec(sqlStr, p.PostID, p.Title, p.Content, p.UserID, p.CommunityID)
	return err
}

func GetPostDetail(p *models.Post) error {
	sqlStr := ` select post_id, user_id, community_id, title, content, create_time
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

func GetPostList(page, size int) ([]*models.Post, error) {
	var posts []*models.Post
	sqlStr := ` select post_id, user_id, community_id, title, content, create_time
	from post
	limit ?, ?
  `
	rows, err := global.DB.Query(sqlStr, page-1, size)
	if err != nil {
		return nil, err
	}
	// 确保关闭数据库查询结果
	defer rows.Close()
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.PostID, &post.UserID, &post.CommunityID, &post.Title, &post.Content, &post.CreateTime)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil { // 检查循环中的错误
		return nil, err
	}
	return posts, nil

}
