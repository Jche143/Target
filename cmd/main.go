package main

import (
	"Target/conf"
	"Target/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := conf.InitDB()

	defer db.Close()

	r := gin.Default()

	routes.CollectRoutes(r)

	r.Run() // 启动服务，并监听 8080 端口
}
