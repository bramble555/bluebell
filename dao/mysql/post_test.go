package mysql

import (
	"bluebell/models"
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

var db *sql.DB

func init() {
	var err error
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		"root",
		"123456",
		"127.0.0.1",
		"3306",
		"bluebell",
	)
	// 不会连接数据库，不会进行验证参数。只是把连接到数据库的struct给设置了
	// 真正的连接是被需要的时候才进行懒设置的
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		panic("数据库连接失败")
	}
	err = db.Ping()
	if err != nil {
		panic("数据库连接失败")
	}
	// context.Context这个类型可以携带截止时间，取消信号
	// ctx := context.Background() // 此函数连接的时候，不会被取消，也没有截止时间
	fmt.Println("Connected")
}
func TestGetPostList2(t *testing.T) {
	idList := []string{"48656817492332544", "48656850105143296"}
	var posts []*models.Post
	sqlStr := `select post_id, user_id, community_id, title, content, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id,?) 
	limit 0,10
	`
	idStr := strings.Join(idList, ",")
	formattedIds := make([]string, len(idList))

	for i, id := range idList {
		formattedIds[i] = fmt.Sprintf(`"%s"`, id)
	}
	// 使用逗号拼接所有元素
	inOrder := strings.Join(formattedIds, ",")
	t.Logf("inOder为%s", inOrder)
	rows, _ := db.Query(sqlStr, inOrder, idStr)
	defer rows.Close()
	for rows.Next() {
		p := &models.Post{}
		err := rows.Scan(&p.PostID, &p.UserID, &p.CommunityID, &p.Title, &p.Content, &p.CreateTime)
		if err != nil {
			t.Log("err是", err)
			return
		}
		t.Logf("post为%+v", p)
		posts = append(posts, p)
	}
	t.Logf("GetPostList2posts为%+v", posts)
}

func TestGetPostList(t *testing.T) {

	idList := []string{"48656817492332544", "48656850105143296"}
	// 创建一个切片用于存储带有双引号的元素
	formattedIds := make([]string, len(idList))

	for i, id := range idList {
		formattedIds[i] = fmt.Sprintf(`"%s"`, id)
	}
	// 使用逗号拼接所有元素
	inOrder := strings.Join(formattedIds, ",")
	t.Logf("inOder为%s", inOrder)
	var posts []*models.Post
	sqlStr := `select post_id, user_id, community_id, title, content, create_time
	from post
	where post_id in (?)
	order by post_id 
	limit 0,10
	`
	// idStr := strings.Join(idList, ",")
	// t.Log("idstr为", idList)
	rows, _ := db.Query(sqlStr, inOrder)
	defer rows.Close()
	for rows.Next() {
		p := &models.Post{}
		err := rows.Scan(&p.PostID, &p.UserID, &p.CommunityID, &p.Title, &p.Content, &p.CreateTime)
		if err != nil {
			t.Log("err是", err)
			return
		}
		t.Logf("post为%+v", p)
		posts = append(posts, p)
	}
	t.Logf("GetPostList posts为%+v", posts)
}
