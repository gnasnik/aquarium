package routers

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func JwtPayloadFunc(data interface{}) jwt.MapClaims {
	return jwt.MapClaims{}
}

func JwtIdentityHandler(ctx *gin.Context) interface{} {
	return nil
}

func JwtAuthenticatorForUser(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}

func JwtAuthorizatorForUser(data interface{}, ctx *gin.Context) bool {
	return false
}

func JwtUnauthorized(ctx *gin.Context, a int, b string) {}

func JwtUserLoginResponse(ctx *gin.Context, a int, b string, c time.Time) {}

func JwtUserRefreshResponse(ctx *gin.Context, a int, b string, c time.Time) {}

func JwtUserHTTPStatusMessageFunc(e error, ctx *gin.Context) string {
	return ""
}
