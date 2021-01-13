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

func ListJobHandler(c *gin.Context) {
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

	total, jobs, err := sdk.ListJob(ctx, ids, size, page, order)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.ListJobFailed))
		return
	}

	for i, trader := range jobs {
		if _, ok := traderx.Executor[trader.ID]; !ok {
			continue
		}
		jobs[i].Status = traderx.Executor[trader.ID].Status
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{
		"total": total,
		"jobs":  jobs,
	}))

}

func PutJobHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

	ctx := context.Background()
	_, err := sdk.GetUserByID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	var req mod.Job
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errw("parse param failed", "err", err)
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.MissingRequestParams))
		return
	}

	req.UserID = uid
	if req.ID > 0 {
		job, err := sdk.GetJobByID(ctx, req.ID)
		if err != nil {
			c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.JobNotFound))
			return
		}
		job.AlgorithmID = req.AlgorithmID
		job.ExchangeID = req.ExchangeID
		job.Status = req.Status
		job.Description = req.Description
		if err := sdk.UpdateJob(ctx, job); err != nil {
			c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UpdateJobFailed))
			return
		}

		c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
		return
	}

	if err := sdk.AddJob(ctx, &req); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.AddJobFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
}

func DeleteJobHandler(c *gin.Context) {
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

	if err := sdk.DeleteJob(ctx, p.IDs); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.DeleteJobFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
}

func SwitchJobHandler(c *gin.Context) {
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

	job, err := sdk.GetJobByID(ctx, p.ID)
	if err != nil || job == nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.JobNotFound))
		return
	}

	if job.Status != mod.JSRunning && job.Status != mod.JSStop {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.InvalidJob))
		return
	}

	if err := traderx.Switch(p.ID); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.SwitchTraderFailed))
		return
	}

	job.LastRunAt = time.Now()
	if job.Status == mod.JSRunning {
		job.Status = mod.JSStop
	}
	if job.Status == mod.JSStop {
		job.Status = mod.JSRunning
	}

	if err := sdk.UpdateJob(ctx, job); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UpdateJobFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
}
