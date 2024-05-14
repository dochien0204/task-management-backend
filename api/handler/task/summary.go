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
	"time"

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
	id, err := taskService.WithTrx(trxHandle).CreateTask(claims.UserId, taskPayload)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	//Response in JSON
	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
		Results: map[string]interface{}{
			"id": id,
		},
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

	trxHandle := ctx.MustGet("db_trx").(*gorm.DB)
	err = taskService.WithTrx(trxHandle).UpdateTask(payload)
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

func updateTaskStatus(ctx *gin.Context, taskService task.UseCase) {
	var payload payload.TaskStatusUpdatePayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	trxHandle := ctx.MustGet("db_trx").(*gorm.DB)
	err = taskService.WithTrx(trxHandle).UpdateTaskStatus(payload.Id, payload.StatusId)
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

func getListTaskByDate(ctx *gin.Context, taskService task.UseCase) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	timeOffset := util.GetDataFromHeader(ctx, "Time-Offset")
	projectId := ctx.Query("projectId")
	projectIdInt, err := strconv.Atoi(projectId)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	fromDate, _ := time.Parse(config.LAYOUT, from)
	toDate, _ := time.Parse(config.LAYOUT, to)

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

	listTask, err := taskService.GetListTaskByDate(projectIdInt, claims.UserId, timeOffset, fromDate, toDate)
	if err != nil {
		util.HandleException(ctx, http.StatusBadGateway, entity.ErrBadRequest)
		return
	}

	response := presenter.BasicResponse {
		Status: fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
		Results: convertListTaskByDateToPresenter(listTask),
	}

	ctx.JSON(http.StatusOK, response)
}

func createDiscussion(ctx *gin.Context, taskService task.UseCase) {
	var payload payload.DiscussionPayload
	err := ctx.ShouldBindJSON(&payload)
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
	err = taskService.WithTrx(trxHandle).CreateDiscussionTask(claims.UserId, payload.TaskId, payload.Comment)
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

func getListDiscussionByTask(ctx *gin.Context, taskService task.UseCase) {
	taskId := ctx.Query("taskId")
	taskIdConv, _ := strconv.Atoi(taskId)
	page := util.GetPage(ctx, "page")
	pageSize := util.GetPageSize(ctx, "size")
	sortBy := ctx.Query("sortBy")
	sortType := ctx.Query("sortType")

	listDiscussion, count, err := taskService.GetListDiscussionOfTask(taskIdConv, page, pageSize, sortBy, sortType)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	response := presenter.PaginationResponse {
		Status: fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
		Results: convertListDiscussionToPresenter(listDiscussion),
		Pagination: presenter.Pagination {
			Count: count,
			NumPages: int(util.CalculateTotalPages(count, pageSize)),
			DisplayRecord: len(listDiscussion),
			Page: page,
		},
	}

	ctx.JSON(http.StatusOK, response)
}

func getDocumentTaskUrl(ctx *gin.Context) {
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

	document := ctx.Query("document")
	taskId := ctx.Query("taskId")
	taskIdInt, err := strconv.Atoi(taskId)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	urlView, err := util.GeneratePresignViewFileURLS3(taskIdInt, document)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	response := presenter.BasicResponse{
		Status: fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
		Results: map[string]interface{}{
			"url": urlView,
		},
	}

	ctx.JSON(http.StatusOK, response)
}

func getListTaskProjectByUserAndStatus(ctx *gin.Context, taskService task.UseCase) {
	projectId := ctx.Query("projectId")
	projectIdConv, _ := strconv.Atoi(projectId)
	statusId := ctx.Query("statusId")
	statusIdConv, _ := strconv.Atoi(statusId)
	userId := ctx.Query("userId")
	userIdConv, _ := strconv.Atoi(userId)
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

	listTask, count, err := taskService.GetListTaskProjectByUserAndStatus(projectIdConv, userIdConv, statusIdConv, page, pageSize, sortType, sortBy)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, err)
		return
	}

	response := presenter.PaginationResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
		Results: convertListTaskProjectByUserAndStatusToPresenter(listTask),
		Pagination: presenter.Pagination {
			Count: count,
			NumPages: int(util.CalculateTotalPages(count, pageSize)),
			DisplayRecord: len(listTask),
			Page: page,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func deleteTask(ctx *gin.Context, taskService task.UseCase) {
	taskId := ctx.Query("taskId")
	taskIdInt, _ := strconv.Atoi(taskId)

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
	err = taskService.WithTrx(trxHandle).DeleteTask(claims.UserId, taskIdInt)
	if err != nil {
		switch err {
		case entity.ErrNotHavePermissionDeleteTask:
			util.HandleException(ctx, http.StatusForbidden, err)
			return
		default:
			util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
			return
		}
	}

	response := presenter.BasicResponse {
		Status: fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
	}

	ctx.JSON(http.StatusOK, response)
}