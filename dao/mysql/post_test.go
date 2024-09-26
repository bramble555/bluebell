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
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		panic("数据库连接失败")
	}
	err = db.Ping()
	if err != nil {
		panic("数据库连接失败")
	}
	fmt.Println("Connected")
}

func TestGetPostList2(t *testing.T) {
	idList := []string{"48656817492332544", "48656850105143296"}
	var posts []*models.Post

	// 使用逗号拼接所有元素
	idInStr := strings.Join(idList, "','")
	idOrderStr := strings.Join(idList, ",")

	// 构建 SQL 查询
	sqlStr := fmt.Sprintf(`select post_id, user_id, community_id, title, content, create_time
	from post
	where post_id in ('%s')
	order by FIND_IN_SET(post_id,'%s') 
	limit %d,%d
	`, idInStr, idOrderStr, 0, 10)
	// 准备预编译语句
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		t.Fatalf("prepare error: %s", err.Error())
	}
	defer stmt.Close()

	// 执行查询
	rows, err := stmt.Query()
	if err != nil {
		t.Fatalf("query error: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		p := &models.Post{}
		err := rows.Scan(&p.PostID, &p.UserID, &p.CommunityID, &p.Title, &p.Content, &p.CreateTime)
		if err != nil {
			t.Fatalf("scan error: %s", err.Error())
		}
		t.Logf("post为%+v", p)
		posts = append(posts, p)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("error iterating over rows: %s", err.Error())
	}
	t.Logf("posts:%+v", posts)

}
