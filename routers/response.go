package routers

import (
	"aquarium/comm"
	"aquarium/errors"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(data comm.JsonObj) gin.H {
	return gin.H{
		"success": true,
		"data":    data,
	}
}

func ResponseFailWithErrorCode(code errors.ErrorCode) gin.H {
	return gin.H{
		"success": false,
		"code":    code,
		"msg":     errors.GetMsg(code),
	}
}
