package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "登录状态有效",
		"meta":  "400",
		"error": "null",
	})
	return
}
