package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"aquarium/comm"
	"aquarium/errors"
	"aquarium/sdk"
	"aquarium/utils/log"

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
	Guid   string `json:"guid"`
	UserID int64  `json:"user_id"`
}

type userAuthInfo struct {
	CurrToken string `json:"curr_token"`
	LastToken string `json:"last_token"`
}

func JwtPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*authResponse); ok {
		return jwt.MapClaims{
			identityKey: v.Guid,
			"user_id":   v.UserID,
		}
	}
	return jwt.MapClaims{}
}

func JwtIdentityHandler(ctx *gin.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return &authResponse{
		Guid:   claims[identityKey].(string),
		UserID: int64(claims["user_id"].(float64)),
	}
}

func JwtAuthenticatorForUser(ctx *gin.Context) (interface{}, error) {
	var loginVals login
	if err := ctx.ShouldBind(&loginVals); err != nil {
		return "", errors.Error[errors.MissingRequestParams]
	}
	userID := loginVals.UserID
	password := loginVals.Password

	log.Debugw("JwtAuthenticatorForUser", "uid", userID, "type", loginVals.LoginType)
	switch loginVals.LoginType {
	case GuestLogin:
		return GuestAuth(userID)
	case PhoneLogin:
		return PhoneAuth(userID, password, false)
	}

	return nil, errors.Error[errors.UnknownLoginType]
}

func GuestAuth(userid int64) (interface{}, error) {
	// implement me
	return nil, nil
}

func PhoneAuth(userid int64, password string, checkAdmin bool) (interface{}, error) {
	user, err := sdk.GetUserByID(context.Background(), userid)
	if err != nil {
		return nil, errors.Error[errors.UserNotFound]
	}
	log.Info("user %v", user)
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.Error[errors.InvalidPassword]
	}
	// if user.IsBanned {
	// 	return nil, nil
	// }
	// TODO: check user role
	return &authResponse{Guid: user.Guid, UserID: user.ID}, nil
}

func JwtAuthorizatorForUser(data interface{}, ctx *gin.Context) bool {
	// if v, ok := data.(*authResponse); ok && v.UserID == 10000 {
	// 	return true
	// }
	// return false
	return true
}

func JwtUnauthorized(ctx *gin.Context, code int, message string) {
	ResponseFailWithErrorMsg(code, message)
}

func JwtUserLoginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	jToken, err := AuthUserMiddleware.ParseTokenString(token)
	claims := jwt.ExtractClaimsFromToken((jToken))
	userID := int64(claims["user_id"].(float64))

	authInfo := userAuthInfo{CurrToken: token}
	_, err = json.Marshal(&authInfo)
	if err != nil {
		log.Errw("create auth info fail", "id", userID, "err", err)
		ctx.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.TokenCreateFail))
		return
	}

	//TODO: write to redis

	ctx.JSON(code, ResponseSuccess(comm.JsonObj{
		"token":     token,
		"expire":    expire.Format(time.RFC3339),
		"expire_ts": expire.Unix(),
	}))
}

func JwtUserRefreshResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	//TODO: check from redis
	ctx.JSON(code, ResponseSuccess(comm.JsonObj{
		"token":     token,
		"expire":    expire.Format(time.RFC3339),
		"expire_ts": expire.Unix(),
	}))
}

func JwtUserHTTPStatusMessageFunc(e error, ctx *gin.Context) string {
	return e.Error()
}
