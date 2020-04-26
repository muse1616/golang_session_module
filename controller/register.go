package controller

import (
	"demo_gin/model"
	"demo_gin/session"
	"demo_gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册handler
func Register(ctx *gin.Context) {
	var r model.User
	if err := ctx.Bind(&r); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg":   "服务器错误",
			"meta":  "401",
			"error": err.Error(),
		})
		return
	}
	result, err := model.RegisterDao(&r)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg":   "注册失败",
			"meta":  "401",
			"error": err.Error(),
		})
		return
	}
	if result == true {
		sessionId:=utils.GenerateUUid()
		_ = session.SetSession(sessionId, r.Id)
		ctx.SetCookie("session_id",sessionId,60,"/","127.0.0.1",false,true)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":   "注册成功",
			"meta":  "400",
			"error": "none",
		})
		return
	} else {
		 ctx.JSON(http.StatusForbidden, gin.H{
			"msg":   "注册失败",
			"meta":  "401",
			"error": "none",
		})
		return
	}
}
