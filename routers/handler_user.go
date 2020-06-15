package routers

import (
	"aquarium/errors"
	"context"
	"net/http"

	"aquarium/sdk"
	"aquarium/sdk/mod"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateNewUserHandler(c *gin.Context) {
	type post struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
		// Role     string `json:"role"`
	}

	var p post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.MissingRequestParams))
		return
	}

	if p.UserID <= 0  || p.Password = "" {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.InvalidRequestParams))
		return
	}

	passHash ,err := bcrypt.GenerateFromPassword([]byte(p.Password),bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.GeneratePasswordFail))
		return 
	}

	user := &mod.User{
		ID: p.UserID,
		Password: p.Password,
	}

	if err := sdk.CreateUser(context.Background(),user); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.CreateNewUserFail))
		return 
	}

	c.JSON(http.StatusOK, ResponseSuccess(common.JsonObj{
		"user_id": p.UserID,
		"password":  p.Password,
	}))
}
