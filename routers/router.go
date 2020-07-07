package routers

import (
	"github.com/frankffenn/aquarium/comm"
	"github.com/frankffenn/aquarium/config"
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

	gin.SetMode(config.Configs.RunMode)
	r := gin.New()
	if config.Configs.RunMode == gin.DebugMode {
		c := cors.DefaultConfig()
		c.AllowAllOrigins = true
		c.AllowMethods = []string{"GET", "PUT", "POST", "OPTIONS"}
		c.AllowHeaders = []string{"*"}
		c.MaxAge = time.Hour
		r.Use(cors.New(c))
	} else {
		r.Use(options)
	}

	apiV1 := r.Group("/api/v1")

	jwtAuthUserMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:                 "User",
		Key:                   []byte(config.Configs.JwtUserSecret),
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

	usr := apiV1.Group("/user")
	// Temp api for testing
	usr.POST("/new", CreateNewUserHandler)

	usr.POST("/login", AuthUserMiddleware.LoginHandler)
	usr.Use(AuthUserMiddleware.MiddlewareFunc())
	usr.GET("/info", GetUserHandler)
	usr.GET("/list", ListUserHandler)

	exch := apiV1.Group("/exchange")
	exch.GET("/types", TypesExchangeHandler)
	exch.Use(AuthUserMiddleware.MiddlewareFunc())
	exch.GET("/list", ListExchangeHandler)
	exch.POST("/put", PutExchangeHandler)
	exch.POST("/del", DeleteExchangeHandler)

	algo := apiV1.Group("/algorithm")
	algo.Use(AuthUserMiddleware.MiddlewareFunc())
	algo.GET("/list", ListAlgorithmHandler)
	algo.POST("/put", PutAlgorithmHandler)
	algo.POST("/del", DeleteAlgorithmHandler)

	trad := apiV1.Group("/trader")
	trad.Use(AuthUserMiddleware.MiddlewareFunc())
	trad.GET("/list", ListTraderHandler)
	trad.POST("/put", PutTraderHandler)
	trad.POST("/del", DeleteTraderHandler)

	return r
}
