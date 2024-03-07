package routes

import (
	"github.com/gin-gonic/gin"
	"Target/service"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// 主页
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to my website",
		})
	})
	
	// 注册
	r.POST("/register", service.Register)

	// 登录
	r.POST("/login", service.Login)

	return r
}
