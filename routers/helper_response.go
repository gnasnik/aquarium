package routers

import "github.com/gin-gonic/gin"

func ResponseFailWithErrorMsg(code int, msg string) gin.H {
	return gin.H{
		"success": false,
		"code":    code,
		"msg":     msg,
	}
}
