package routers

import (
	"aquarium/comm"
	"aquarium/config"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func options(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
	}
}

var (
	TokenTimeout        = time.Hour * 24
	TokenRefreshTimeout = time.Hour * 24 * 30
	AuthUserMiddleware  *AppJWTMiddleware
)

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

	jwtAuthUserMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:                 "User",
		Key:                   []byte(config.JwtUserSecret),
		Timeout:               TokenTimeout,
		MaxRefresh:            TokenRefreshTimeout,
		IdentityKey:           identityKey,
		PayloadFunc:           JwtPayloadFunc,
		IdentityHandler:       JwtIdentityHandler,
		Authenticator:         JwtAuthenticatorForUser,
		Authorizator:          JwtAuthorizatorForUser,
		Unauthorized:          JwtUnauthorized,
		LoginResponse:         JwtUserLoginResponse,
		RefreshResponse:       JwtUserRefreshResponse,
		HTTPStatusMessageFunc: JwtUserHTTPStatusMessageFunc,
	})

	if err != nil {
		log.Fatal(err)
	}

	AuthUserMiddleware = &AppJWTMiddleware{
		jwtAuthUserMiddleware,
	}

	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{"pong": true}))
	})

	apiV1.POST("/user/login", AuthUserMiddleware.LoginHandler)

	return r
}
