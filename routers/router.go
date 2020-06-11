package routers

import (
	"aquarium/comm"
	"aquarium/config"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func options(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
	}
}

func InitRouter() *gin.Engine {

	gin.SetMode(config.RunMode)
	r := gin.New()
	if config.RunMode == gin.DebugMode {
		c := cors.DefaultConfig()
		c.AllowAllOrigins = true
		c.AllowMethods = []string{"GET", "POST", "OPTION"}
		c.AllowHeaders = []string{"*"}
		c.MaxAge = time.Hour
		r.Use(cors.New(c))
	} else {
		r.Use(options)
	}

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{"pong": true}))
	})

	return r
}
