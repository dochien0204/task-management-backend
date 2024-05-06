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
	"strconv"

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
	err = taskService.WithTrx(trxHandle).CreateTask(claims.UserId, taskPayload)
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

func getListTaskOfProject(ctx *gin.Context, taskService task.UseCase) {
	projectId := ctx.Query("projectId")
	projectIdConv, _ := strconv.Atoi(projectId)
	page := util.GetPage(ctx, "page")
	pageSize := util.GetPageSize(ctx, "size")
	sortBy := ctx.Query("sortBy")
	sortType := ctx.Query("sortType")

	//Get URL param
	token, err := util.GetToken(ctx)
	if err != nil {
		util.HandleException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
		return
	}

	_, err = util.ParseAccessToken(token)
	if err != nil {
		util.HandleException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
		return
	}

	listTask, listStatus, err := taskService.GetListTaskOfProject(projectIdConv, page, pageSize, sortType, sortBy)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, err)
		return
	}

	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
		Results: convertListTaskToPresenter(listTask, listStatus),
	}
	ctx.JSON(http.StatusOK, response)
}

func getTaskDetail(ctx *gin.Context, taskService task.UseCase) {
	taskId := ctx.Query("taskId")
	taskIdInt, _ := strconv.Atoi(taskId)

	task, err := taskService.GetTaskDetail(taskIdInt)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	response := presenter.BasicResponse {
		Status: fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
		Results: convertTaskDetailEntityToPresenter(task),
	}

	ctx.JSON(http.StatusOK, response)
}

func updateTask(ctx *gin.Context, taskService task.UseCase) {
	var payload payload.TaskUpdatePayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	err = taskService.UpdateTask(payload)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	response := presenter.BasicResponse {
		Status: fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
	}

	ctx.JSON(http.StatusOK, response)
}