package model

import "github.com/jinzhu/gorm"

// type User struct {
// 	gorm.Model
// 	Name     string `gorm:"varchar(20);not null"`
// 	Username string `gorm:"varchar(20);not null"`
// 	Password string `gorm:"size:255;not null"`
// }

type User struct {
	gorm.Model
	Name     string `query:"name" form:"name" json:"name" gorm:"varchar(20);not null"`
	Username string `query:"username" form:"username" json:"username" gorm:"varchar(20);not null"`
	Password string `query:"password" form:"password" json:"password" gorm:"varchar(20);not null"`
}