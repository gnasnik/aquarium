package routers

import (
	"context"
	"time"

	"aquarium/sdk"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	LoginType string `form:"login_type" json:"login_type" binding:"required"`
	UserID    int64  `form:"user_id" json:"user_id" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

type authResponse struct {
	Guid   string
	UserID int64
}

func JwtPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*authResponse); ok {
		return jwt.MapClaims{
			identityKey: v.Guid,
			"UserId":    v.UserID,
		}
	}
	return jwt.MapClaims{}
}

func JwtIdentityHandler(ctx *gin.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return &authResponse{
		Guid:   claims[identityKey].(string),
		UserID: claims["UserId"].(int64),
	}
}

func JwtAuthenticatorForUser(ctx *gin.Context) (interface{}, error) {
	var loginVals login
	if err := ctx.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.UserID
	password := loginVals.Password

	switch loginVals.LoginType {
	case GuestLogin:
		return GuestAuth(userID)
	case PhoneLogin:
		return PhoneAuth(userID, password, false)
	}

	return nil, jwt.ErrFailedAuthentication
}

func GuestAuth(userid string) (interface{}, error) {
	// implement me
	return nil, nil
}

func PhoneAuth(userid int64, password string, checkAdmin bool) (interface{}, error) {
	user, err := sdk.GetUserByID(context.Background(), userid)
	if err != nil {
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	if user.IsBanned {
		return nil, nil
	}
	// TODO: check user role
	return &authResponse{Guid: user.Guid, UserID: user.ID}, nil
}

func JwtAuthorizatorForUser(data interface{}, ctx *gin.Context) bool {
	if v, ok := data.(*authResponse); ok && v.UserID == 10000 {
		return true
	}
	return false
}

func JwtUnauthorized(ctx *gin.Context, code int, message string) {
	ResponseFailWithErrorMsg(code, message)
}

func JwtUserLoginResponse(ctx *gin.Context, a int, b string, c time.Time) {}

func JwtUserRefreshResponse(ctx *gin.Context, a int, b string, c time.Time) {}

func JwtUserHTTPStatusMessageFunc(e error, ctx *gin.Context) string {
	return e.Error()
}
