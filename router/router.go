package router

import (
	"demo_gin/controller"
	"demo_gin/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	//默认路由
	r := gin.Default()
	//v1组路由
	v1Group := r.Group("v1")
	{
		//注册 注册完成 cookie保存60秒
		v1Group.POST("/register", controller.Register)
		//登录
		v1Group.POST("/login", controller.Login)
		//需要登录状态授权的接口
		v1Group.GET("/auth", middleware.IsLogin(),controller.Auth)
	}
	//开启路由
	err := r.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	//返回路由实例
	return r
}
