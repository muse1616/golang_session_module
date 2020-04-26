package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

//初始化mysql
func InitMysql() (err error) {
	dsn := "root:19991116hjj@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//测试连通
	err = DB.DB().Ping()
	if err != nil {
		return err
	}
	return
}

// 关闭mysql连接
func CloseMysqlConnection() (err error) {
	err = DB.Close()
	return
}
