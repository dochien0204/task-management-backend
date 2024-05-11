package handler

import (
	"fmt"
	"net/http"
	projectPayload "source-base-go/api/payload/project"
	"source-base-go/api/presenter"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/project"
	"strconv"
	"time"

	"github.com/gin-contrib/i18n"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createProject(ctx *gin.Context, projectService project.UseCase) {
	//Get URL param
	var payload projectPayload.ProjectPayload
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
	err = projectService.WithTrx(trxHandle).CreateProject(claims.UserId, convertProjectPayloadToEntity(payload))
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, err)
		return
	}

	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
	}
	ctx.JSON(http.StatusOK, response)
}

func getListProjectOfUser(ctx *gin.Context, projectService project.UseCase) {
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

	claims, err := util.ParseAccessToken(token)
	if err != nil {
		util.HandleException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
		return
	}

	listProject, err := projectService.GetListProjectOfUser(claims.UserId, page, pageSize, sortType, sortBy)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, err)
		return
	}

	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
		Results: convertListProjectEntityToPresenter(listProject),
	}
	ctx.JSON(http.StatusOK, response)
}

func addListMemberToProject(ctx *gin.Context, projectService project.UseCase) {
	//Get URL param
	var payload projectPayload.ListUserIdProjectPayload
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
	err = projectService.WithTrx(trxHandle).AddListMemberToProject(claims.UserId, payload.ProjectId, payload.RoleId, payload.ListUserId)
	if err != nil {
		switch err {
		case entity.ErrForbidden:
			util.HandleException(ctx, http.StatusForbidden, err)
			return
		default:
			util.HandleException(ctx, http.StatusBadRequest, err)
			return
		}
	}

	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
	}
	ctx.JSON(http.StatusOK, response)
}

func getProjectDetail(ctx *gin.Context, projectService project.UseCase) {
	projectId := ctx.Query("projectId")
	projectIdConv, err := strconv.Atoi(projectId)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	projectDetail, err := projectService.GetProjectDetail(projectIdConv)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
		Results: convertProjectDetailToPresenter(projectDetail),
	}
	ctx.JSON(http.StatusOK, response)
}

func getListMemberTaskCount(ctx *gin.Context, projectService project.UseCase) {
	page := util.GetPage(ctx, "page")
	pageSize := util.GetPageSize(ctx, "size")
	sortBy := ctx.Query("sortBy")
	sortType := ctx.Query("sortType")
	projectId := ctx.Query("projectId")
	projectIdInt, err := strconv.Atoi(projectId)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	listTask, count, err := projectService.GetListMemberByProject(projectIdInt, page, pageSize, sortType, sortBy)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, err)
		return
	}

	response := presenter.PaginationResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
		Results: convertListMemberTaskCountToPresenter(listTask),
		Pagination: presenter.Pagination {
			Count: count,
			NumPages: int(util.CalculateTotalPages(count, pageSize)),
			DisplayRecord: len(listTask),
			Page: page,
		},
	}
	ctx.JSON(http.StatusOK, response)
}

func getListActivityProjectByDate(ctx *gin.Context, projectService project.UseCase) {
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

	_, err = util.ParseAccessToken(token)
	if err != nil {
		util.HandleException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
		return
	}

	listActivity, err := projectService.GetListActivityProjectByDate(projectIdInt, timeOffset, fromDate, toDate)
	if err != nil {
		util.HandleException(ctx, http.StatusBadGateway, entity.ErrBadRequest)
		return
	}

	response := presenter.BasicResponse {
		Status: fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
		Results: convertListActivityProjectByDateToPresenter(listActivity),
	}

	ctx.JSON(http.StatusOK, response)
}

func getUserOverviewTaskProject(ctx *gin.Context, projectService project.UseCase) {
	projectId := ctx.Query("projectId")
	UserId := ctx.Query("userId")
	projectIdInt, _ := strconv.Atoi(projectId)
	userIdInt, _ := strconv.Atoi(UserId)

	userOpenTaskCount, userClosedTaskCount, userProjectRole, err := projectService.GetOverviewUserTaskProject(projectIdInt, userIdInt)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	response := presenter.BasicResponse {
		Status: fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
		Results: convertUserProjectOverviewToPresenter(userOpenTaskCount, userClosedTaskCount, userProjectRole),
	}

	ctx.JSON(http.StatusOK, response)
}
