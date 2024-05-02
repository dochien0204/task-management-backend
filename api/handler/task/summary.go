package handler

import (
	"fmt"
	"net/http"
	payload "source-base-go/api/payload/task"
	"source-base-go/api/presenter"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/task"

	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createTask(ctx *gin.Context, taskService task.UseCase) {
	var taskPayload payload.TaskPayload
	err := ctx.ShouldBindJSON(&taskPayload)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	token, err := util.GetToken(ctx)
	if err != nil {
		util.HandleException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
		return
	}

	claims, err := util.ParseAccessToken(token)
	if err != nil {
		util.HandleException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
		return
	}

	trxHandle := ctx.MustGet("db_trx").(*gorm.DB)
	err = taskService.WithTrx(trxHandle).CreateTask(claims.UserId, &taskPayload)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	//Response in JSON
	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
	}

	ctx.JSON(http.StatusOK, response)
}