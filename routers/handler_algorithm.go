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

func ListAlgorithmHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

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
	total, Algorithms, err := sdk.ListAlgorithm(ctx, uid, size, page, order)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.ListAlgorithmFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{
		"total":      total,
		"algorithms": Algorithms,
	}))

}

func PutAlgorithmHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

	ctx := context.Background()
	_, err := sdk.GetUserByID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	var req mod.Algorithm
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errw("parse param failed", "err", err)
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.MissingRequestParams))
		return
	}

	if req.ID > 0 {
		algorithm, err := sdk.GetAlgorithmByID(ctx, req.ID)
		if err != nil {
			c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.AlgorithmNotFound))
			return
		}
		algorithm.Name = req.Name
		algorithm.Description = req.Description
		algorithm.Script = req.Script
		algorithm.EvnDefault = req.EvnDefault
		if err := sdk.UpdateAlgorithm(ctx, algorithm); err != nil {
			c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UpdateAlgorithmFailed))
			return
		}

		c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
		return
	}

	if err := sdk.AddAlgorithm(ctx, &req); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.AddAlgorithmFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
}

func DeleteAlgorithmHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

	ctx := context.Background()
	_, err := sdk.GetUserByID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	type post struct {
		IDs []int64 `json:"ids"`
	}

	var p post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.MissingRequestParams))
		return
	}

	if err := sdk.DeleteAlgorithm(ctx, p.IDs); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.DeleteAlgorithmFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
}
