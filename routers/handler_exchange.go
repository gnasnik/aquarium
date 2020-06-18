package routers

import (
	"aquarium/comm"
	"aquarium/errors"
	"aquarium/sdk"
	"aquarium/sdk/mod"
	"aquarium/utils/log"
	"context"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func ListExchangeHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))
	level := int64(claims["level"].(float64))

	page := com.StrTo(c.Query("page")).MustInt64()
	size := com.StrTo(c.Query("size")).MustInt64()
	order := c.Query("order")

	if page <= 0 {
		page = 1
	}

	if size <= 0 {
		size = 20
	}

	ctx := context.Background()
	_, users, err := sdk.ListUser(ctx, uid, level, -1, 1, order)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.ListUserFailed))
		return
	}

	var ids []interface{}
	for _, x := range users {
		ids = append(ids, x.ID)
	}

	total, exchanges, err := sdk.ListExchange(ctx, ids, size, page, order)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.ListExchangeFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{
		"total":     total,
		"exchanges": exchanges,
	}))

}

func PutExchangeHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

	ctx := context.Background()
	_, err := sdk.GetUserByID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	var req mod.Exchange
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Err("parse param failed", "err", err)
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.MissingRequestParams))
		return
	}

	if req.ID > 0 {
		exchange, err := sdk.GetExchangeByID(ctx, req.ID)
		if err != nil {
			c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.ExchangeNotFound))
			return
		}
		exchange.Name = req.Name
		exchange.Type = req.Type
		exchange.AccessKey = req.AccessKey
		exchange.SecretKey = req.SecretKey
		if err := sdk.UpdateExchange(ctx, exchange); err != nil {
			c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UpdateExchangeFailed))
			return
		}

		c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
		return
	}

	if err := sdk.AddExchange(ctx, &req); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.AddExchangeFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
}
