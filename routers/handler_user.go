package routers

import (
	"aquarium/errors"
	"context"
	"net/http"

	"aquarium/sdk"
	"aquarium/sdk/mod"

	"aquarium/comm"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
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
