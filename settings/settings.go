package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了")
	})
	return
}
