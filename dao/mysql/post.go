package mysql

import (
	"bluebell/global"
	"bluebell/models"
	"database/sql"
	"fmt"
	"strings"
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
func CheckPIDExist(pid int) (bool, error) {
	var temp int
	sqlStr := `select post_id from post where post_id = ?`
	err := global.DB.QueryRow(sqlStr, pid).Scan(&temp)
	if err != nil {
		if err == sql.ErrNoRows { // 或者根据您的数据库驱动检查相应的“没有找到行”的错误
			return false, nil // 或者返回特定的错误，例如 errors.New("post not found")
		}
		return false, err // 返回其他类型的错误
	}
	return true, nil
}
func GetPostList(page, size int) ([]*models.Post, error) {
	var posts []*models.Post
	sqlStr := ` select post_id, user_id, community_id, title, content, create_time
	from post
	ORDER BY create_time DESC
	limit ?, ?
 	 `
	rows, err := global.DB.Query(sqlStr, (page-1)*size, size)
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
func GetPostList2(idList []string, page, size int) ([]*models.Post, error) {
	var posts []*models.Post
	if len(idList) == 0 {
		return nil, nil
	}
	// 使用逗号拼接所有元素
	idInStr := strings.Join(idList, "','")
	idOrderStr := strings.Join(idList, ",")

	// 构建 SQL 查询
	sqlStr := fmt.Sprintf(`select post_id, user_id, community_id, title, content, create_time
	from post
	where post_id in ('%s')
	order by FIND_IN_SET(post_id,'%s') 
	limit %d,%d
	`, idInStr, idOrderStr, (page-1)*size, size)
	stmt, err := global.DB.Prepare(sqlStr)
	if err != nil {
		global.Log.Errorln("db prepare error", err.Error())
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		global.Log.Errorln("db query error", err.Error())
		return nil, err
	}
	global.Log.Debugln("idstr为", idOrderStr, "页数为", page, "数量为", size, "(page-1)*size为", (page-1)*size)
	defer rows.Close()
	for rows.Next() {
		p := &models.Post{}
		err := rows.Scan(&p.PostID, &p.UserID, &p.CommunityID, &p.Title, &p.Content, &p.CreateTime)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
