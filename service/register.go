package service

import (
	model "Target/model"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// 数据验证
func CheckInfo(c *gin.Context, name, id, passw string) bool {
	if len(id) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名不能为空",
		})
		return false
	}

	if len(name) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "昵称不能为空",
		})
		return false
	}

	if len(passw) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return false
	}

	// 判断用户是否存在
	// 待添加数据库

	return true
}

func Register(c *gin.Context, name, id, passw string) {
	// 数据验证
	if !CheckInfo(c, name, id, passw) {
		return
	}

	// 用户注册
	newUser := model.User{
		Name:     name,
		ID:       id,
		Password: passw,
	}

	// 先保存到本地
	f, err := os.Open("user.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(newUser.ID + " " + newUser.Password + " " + newUser.Name + "\n")
	if err != nil {
		log.Fatal(err)
	}

}
