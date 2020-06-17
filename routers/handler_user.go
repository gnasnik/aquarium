package routers

import (
	"aquarium/errors"
	"aquarium/utils/log"
	"context"
	"net/http"

	"aquarium/sdk"
	"aquarium/sdk/mod"

	"aquarium/comm"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"github.com/unknwon/com"
	"golang.org/x/crypto/bcrypt"
)

func CreateNewUserHandler(c *gin.Context) {
	type post struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var p post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.MissingRequestParams))
		return
	}

	if p.Username == "" || p.Password == "" {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.InvalidRequestParams))
		return
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.GeneratePasswordFail))
		return
	}

	user := &mod.User{
		// ID:       p.UserID,
		Username: p.Username,
		Guid:     ksuid.New().String(),
		Password: string(passHash),
	}

	if err := sdk.CreateUser(context.Background(), user); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.CreateNewUserFail))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{
		"user_id":  user.ID,
		"guid":     user.Guid,
		"username": user.Username,
		// "password": p.Password,
	}))
}

func GetUserHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

	log.Debugw("GetUserHandler", "uid", uid)
	// username := c.Query("username")
	user, err := sdk.GetUserByID(context.Background(), uid)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{
		"user": user.ToPlain(),
	}))
}

func ListUserHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))
	size := com.StrTo(c.Query("size")).MustInt64()
	page := com.StrTo(c.Query("page")).MustInt64()
	order := c.Query("order")

	if size <= 0 {
		size = 20
	}

	if page <= 0 {
		page = 1
	}

	user, err := sdk.GetUserByID(context.Background(), uid)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	total, users, err := sdk.ListUser(context.Background(),
		user.ID, user.Level, size, page, order)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	var out []*mod.PlainUser
	for _, x := range users {
		out = append(out, x.ToPlain())
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{
		"total": total,
		"users": out,
	}))
}
