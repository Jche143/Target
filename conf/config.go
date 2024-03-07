package conf

import (
	"Target/model"
	"fmt"

	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

var (
	APPMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

var DB *gorm.DB

// 加载配置文件
func LoadServer(file *ini.File) {
	APPMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func Loadmysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

// 读取配置文件
func Init() {
	// 从本地环境读取环境变量
	// 使用ini.Load加载配置文件
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}
	LoadServer(file)
	Loadmysql(file)
}

func InitDB() *gorm.DB {
	// 初始化数据库连接
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		DbUser,
		DbPassWord,
		DbHost,
		DbPort,
		DbName,
		"utf8",
	)

	db, err := gorm.Open(Db, args)
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