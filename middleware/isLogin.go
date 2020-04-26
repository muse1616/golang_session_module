package middleware

import (
	"demo_gin/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//取cookie
		id, err1 := c.Cookie("session_id")
		userId, err2 := session.GetSession(id)
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"msg":   "登录状态失效",
				"meta":  "401",
				"error": "null",
			})
			c.Abort()
		}

		if userId == id {
			c.Next()
		}

	}
}
