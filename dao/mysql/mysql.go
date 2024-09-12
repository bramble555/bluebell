package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"bluebell/global"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func Init() (db *sql.DB, err error) {
	// 连接字符串，也就是数据源
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)

	// 不会连接数据库，不会进行验证参数。只是把连接到数据库的struct给设置了
	// 真正的连接是被需要的时候才进行懒设置的
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		global.Log.Fatalln(err)
		return
	}

	// context.Context这个类型可以携带截止时间，取消信号
	ctx := context.Background() // 此函数连接的时候，不会被取消，也没有截止时间
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err)
		return
	}
	global.Log.Println("Connected")
	return
}
