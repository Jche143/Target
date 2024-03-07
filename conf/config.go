package conf

import (
	"Target/model"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	driverName string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	driverName = "mysql"
	DbHost := "127.0.0.1"
	DbPort := "3306"
	DbUser := "root"
	DbPassWord := "123456"
	DbName := "db_mall"

	// 初始化数据库连接
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		DbUser,
		DbPassWord,
		DbHost,
		DbPort,
		DbName,
		"utf8",
	)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	// 自动迁移
	db.AutoMigrate(&model.User{})

	DB = db

	return db
}

func GetDB() *gorm.DB {
	return DB
}
