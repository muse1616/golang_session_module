package model

import (
	"demo_gin/dao"
)

func LoginDao(user User) (result bool) {
	//处理登录请求
	dao.DB.SingularTable(true)
	dao.DB.AutoMigrate(&User{})
	var u User
	dao.DB.Where("id = ? AND password = ?", user.Id, user.Password).First(&u)
	if u.Id == user.Id && u.Password == user.Password {
		return true
	}
	return false
}
