package main

import (
	"Target/conf"
	"Target/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	db := conf.InitDB()

	defer db.Close()

	r := gin.Default()

	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to my website",
		})
	})

	// 注册
	r.POST("/register", service.Register)

	// 登录
	r.POST("/login", service.Login)

	r.Run() // 启动服务，并监听 8080 端口
}
