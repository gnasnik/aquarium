package routers

import (
	"context"
	"net/http"

	"github.com/frankffenn/aquarium/comm"
	"github.com/frankffenn/aquarium/errors"
	"github.com/frankffenn/aquarium/sdk"
	"github.com/frankffenn/aquarium/sdk/mod"
	"github.com/frankffenn/aquarium/utils/log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func ListJobLogHandler(c *gin.Context) {
	jobId := com.StrTo(c.Query("job_id")).MustInt64()
	page := com.StrTo(c.Query("page")).MustInt64()
	size := com.StrTo(c.Query("size")).MustInt64()
	order := c.Query("order")

	if page <= 0 {
		page = 1
	}

	if size <= 0 {
		size = 20
	}

	if order == "" {
		order = "created_at"
	}

	total, logs, err := sdk.ListJobLog(context.Background(), jobId, size, page, order)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.ListJobFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{
		"total":   total,
		"jobLogs": logs,
	}))

}

func PutJobLogHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uid := int64(claims["user_id"].(float64))

	ctx := context.Background()
	_, err := sdk.GetUserByID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.UserNotFound))
		return
	}

	var req mod.JobLog
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errw("parse param failed", "err", err)
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.MissingRequestParams))
		return
	}

	req.UserID = uid
	if err := sdk.AddJobLog(ctx, &req); err != nil {
		c.JSON(http.StatusOK, ResponseFailWithErrorCode(errors.AddJobFailed))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(comm.JsonObj{}))
}
