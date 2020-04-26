package controller

import (
	"demo_gin/model"
	"demo_gin/session"
	"demo_gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	var u model.User
	if err := ctx.Bind(&u); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg":   "服务器错误",
			"meta":  "401",
			"error": err.Error(),
		})
		return
	}
	result := model.LoginDao(u)
	if result == true {
		//设置session_id
		sessionId := utils.GenerateUUid()
		_ = session.SetSession(sessionId, u.Id)
		ctx.SetCookie("session_id", sessionId, 60, "/", "127.0.0.1", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":   "登录成功",
			"meta":  "400",
			"error": "null",
		})
		return
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg":   "登录失败",
			"meta":  "401",
			"error": "null",
		})
		return
	}

}
