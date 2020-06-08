package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func options(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
	}
}


func InitRouter() *gin.Engine {
	var err error
	gin.SetMode(config.RunMode)
	r := gin.New()
	if config.RunMode == gin.DebugMode {
		c := cors.DefaultConfig()
		c.AllowAllOrigins = []string{"*"}
		c.AllowMethods = []{"GET","POST","OPTION"}
		c.AllowHeader = []string{"*"}
		c.MaxAge = time.Hour
		r.Use(cors.New(c))
	}else{
		r.Use(options)
	}

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, ResponseSuccess(common.JsonObj{"pong": true}))
	})

	return r
}
