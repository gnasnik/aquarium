package routers

import (
	"aquarium/comm"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(data comm.JsonObj) gin.H {
	return gin.H{
		"success": true,
		"data":    data,
	}
}
