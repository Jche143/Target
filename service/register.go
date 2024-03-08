package service

import (
	"Target/conf"
	model "Target/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 数据验证
func CheckRegisterInfo(c *gin.Context, name, id, passw string, db *gorm.DB) bool {
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

	return true
}

func Register(c *gin.Context) {

	db := conf.GetDB()

	// 获取参数
	var requestUser model.User
	// 使用bind绑定数据
	c.Bind(&requestUser)
	name := requestUser.Name
	username := requestUser.Username
	passw := requestUser.Password

	// 数据验证
	if !CheckRegisterInfo(c, name, username, passw, db) {
		return
	}

	// 判断用户是否存在
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户已存在",
		})
		return
	}

	// 用户注册
	newUser := model.User{
		Name:     name,
		Username: username,
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


func CheckLoginInfo(c *gin.Context, username, passw string, db *gorm.DB) bool {
    if len(username) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名不能为空",
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

	return true
}

func Login(c *gin.Context) {
	db := conf.GetDB()

	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	username := requestUser.Username
	passw := requestUser.Password

	// 数据验证
	if !CheckLoginInfo(c, username, passw, db) {
		return
	}

	// 判断用户是否存在
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户不存在",
		})

		return
	}

	//判断密码是否正确
	if passw != user.Password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})

		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
	})
}
