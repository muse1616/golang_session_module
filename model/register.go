package model

import "demo_gin/dao"



func RegisterDao(user *User) (result bool, err error) {
	//处理注册请求
	dao.DB.SingularTable(true)
	dao.DB.AutoMigrate(&User{})
	if err = dao.DB.Debug().Create(&user).Error; err != nil {
		return
	}
	result = true
	return
}
