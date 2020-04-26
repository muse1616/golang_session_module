package main

import (
	"demo_gin/dao"
	"demo_gin/router"
	"log"
)

func main() {
	//初始化mysql数据库
	if err := dao.InitMysql(); err != nil {
		panic(err)
	}
	log.Println("mysql connect successfully")

	//初始化redis连接池
	dao.InitRedis("127.0.0.1:6379","")
	log.Println("redis pool connect successfully")


	//开启路由
	router.SetUpRouter()
	log.Println("server starts at 127.0.0.1:8080")
}
