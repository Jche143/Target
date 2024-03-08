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

	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to my website",
		})
	})

	r.POST("/test", func(c *gin.Context) {
		var users model.User
		if err := c.Bind(&users); err != nil { //根据req的content type 自动推断如何绑定,form/json/xml等格式
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// username := c.Query("username")
		// password := c.Query("password")

		// username_p := c.PostForm("username")
		// password_p := c.PostForm("password")

		c.JSON(200, gin.H{
			// "username": username,
			// "password": password,
			// "username_p": username_p,
			// "password_p": password_p,
			"username": users.Username,
			"password": users.Password,
		})
	})

	// 注册
	r.POST("/register", service.Register)

	// 登录
	r.POST("/login", service.Login)

	r.Run() // 启动服务，并监听 8080 端口
}
