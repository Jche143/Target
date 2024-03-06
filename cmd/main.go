package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 定义一个路径为 /ping 的 GET 格式路由，并返回 JSON 数据
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin by quanxiaoha.com !",
		})
	})

	r.POST("/login", func(c *gin.Context) {

		// 获取参数
		name := c.PostForm("name")
		id := c.PostForm("id")
		passw := c.PostForm("passw")

		fmt.Println(name, id, passw)
	})

	r.Run() // 启动服务，并监听 8080 端口
}
