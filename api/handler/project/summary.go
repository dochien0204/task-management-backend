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
