package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Run() // 启动服务，并监听 8080 端口
}
