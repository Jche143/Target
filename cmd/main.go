package main

import (
	"Target/conf"
	"Target/model"
	"Target/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	db := conf.InitDB()

	defer db.Close()

	r := gin.Default()

	r.GET("/api/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to my website",
		})
	})

	r.POST("/api/test", func(c *gin.Context) {
		users := model.User{}
		c.Bind(&users)
		username := users.Username
		password := users.Password
		// username := c.Query("username")
		// password := c.Query("password")

		// username_p := c.PostForm("username")
		// password_p := c.PostForm("password")

		c.JSON(200, gin.H{
			// "username": username,
			// "password": password,
			// "username_p": username_p,
			// "password_p": password_p,
			"username":   username,
			"password":   password,
		})
	})

	// 注册
	r.POST("/api/register", service.Register)

	// 登录
	r.POST("/api/login", service.Login)

	r.Run() // 启动服务，并监听 8080 端口
}
