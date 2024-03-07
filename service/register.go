package service

import (
	"Target/conf"
	model "Target/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 数据验证
func CheckInfo(c *gin.Context, name, id, passw string, db *gorm.DB) bool {
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
	var user model.User
	db.Where("id = ?", id).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户已存在",
		})
		return false
	}

	return true
}

func Register(c *gin.Context) {

	db := conf.GetDB()

	// 获取参数
	var requestUser model.User
	// 使用bind绑定数据
	c.Bind(&requestUser)
	name := requestUser.Name
	id := requestUser.Id
	passw := requestUser.Password

	// 数据验证
	if !CheckInfo(c, name, id, passw, db) {
		return
	}

	// 用户注册
	newUser := model.User{
		Name:     name,
		Id:       id,
		Password: passw,
	}

	// 先保存到本地
	// f, err := os.Open("user.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// _, err = f.WriteString(newUser.Id + " " + newUser.Password + " " + newUser.Name + "\n")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db.Create(&newUser)

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
}