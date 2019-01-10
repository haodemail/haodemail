package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

// DB GORM glob handler
var DB *gorm.DB
var Config ServerConfig

// InitDB init database when start application
func init() {
	var err error
	Config = InitConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Config.MySQL.Username,
		Config.MySQL.Password,
		Config.MySQL.Host,
		Config.MySQL.Port,
		Config.MySQL.Database,
	)
	DB, err = gorm.Open("mysql", dsn)
	if err == nil {
		DB.LogMode(true)
		//db.AutoMigrate(&VirtualUser{})
		// 添加外键
		//db.Model(&PublicMailListLog{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		//db.Model(&SystemMailAddress{}).AddForeignKey("system_maillist_id", "system_mail_lists(id)", "RESTRICT", "RESTRICT")
		//db.Model(&Task{}).AddForeignKey("author_id", "users(id)", "RESTRICT", "RESTRICT")

	} else {
		log.Fatal("connect mysql database failed", err)
	}
}
