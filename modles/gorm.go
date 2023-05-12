package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func init() {

	//配置Mysql连接参数
	userName := "root"   // 账号
	passWord := "123456" // 密码
	host := "127.0.0.1"  // 数据库地址，可以是ip可以是域名
	port := "3306"       //数据库端口号
	Dbname := "ginchat"  // 数据库名   "root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, passWord, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("数据库连接失败！", err)
	}
	DB = db

}