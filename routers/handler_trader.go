package routers

import (
	"context"
	"net/http"
	"time"

	"github.com/frankffenn/aquarium/comm"
	"github.com/frankffenn/aquarium/errors"
	"github.com/frankffenn/aquarium/sdk"
	"github.com/frankffenn/aquarium/sdk/mod"
	traderx "github.com/frankffenn/aquarium/trader"
	"github.com/frankffenn/aquarium/utils/log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func ListTraderHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

	algorithmID := com.StrTo(c.Query("algorithmId")).MustInt64()
	ctx := context.Background()

	Traders, err := sdk.ListTrader(ctx, uid, algorithmID)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.ListTraderFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{
		"traders": Traders,
	}))

}

func PutTraderHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

	ctx := context.Background()
	user, err := sdk.GetUserByID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	var req mod.Trader
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errw("parse param failed", "err", err)
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.MissingRequestParams))
		return
	}

	req.UserID = uid
	if req.ID > 0 {
		if err := sdk.UpdateTrader(ctx, &req); err != nil {
			c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UpdateTraderFailed))
			return
		}

		c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
		return
	}

	req.UserID = user.ID
	if err := sdk.AddTrader(ctx, &req); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.AddTraderFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
}

func DeleteTraderHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

	ctx := context.Background()
	_, err := sdk.GetUserByID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	type post struct {
		ID int64 `json:"id"`
	}

	var p post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.MissingRequestParams))
		return
	}

	trader, err := sdk.GetTraderByID(ctx, p.ID)
	if err != nil || trader == nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.TraderNotFound))
		return
	}

	if err := sdk.DeleteTrader(ctx, []int64{p.ID}); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.DeleteTraderFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
}

func SwitchTraderHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

	ctx := context.Background()
	_, err := sdk.GetUserByID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	type post struct {
		ID int64 `json:"id"`
	}

	var p post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.MissingRequestParams))
		return
	}

	trader, err := sdk.GetTraderByID(ctx, p.ID)
	if err != nil || trader == nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.TraderNotFound))
		return
	}

	if err := traderx.Switch(p.ID); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.SwitchTraderFailed))
		return
	}

	trader.LastRunAt = time.Now()
	if err := sdk.UpdateTrader(ctx, trader); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UpdateTraderFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
}
