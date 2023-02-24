package conf

import (
	"fmt"
	"github.com/gin-demo/go-project/model"
	"github.com/spf13/viper"
)

var (
	User   string
	Pass   string
	IP     string
	Port   string
	DBName string
)

func ConfigInit() {
	// 使用viper读取配置文件
	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./conf/") // path to look for the config file in
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// 初始化MySql
	ReadMySqlConf()
	dsn := User + ":" + Pass + "@tcp(" + IP + ":" + Port + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	model.MySQLInit(dsn)
}

func ReadMySqlConf() {
	User = viper.GetString("MySql.User")
	Pass = viper.GetString("MySql.Pass")
	IP = viper.GetString("MySql.IP")
	Port = viper.GetString("MySql.Port")
	DBName = viper.GetString("MySql.Database")
}
