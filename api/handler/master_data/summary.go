package handler

import (
	"fmt"
	"net/http"
	"source-base-go/api/presenter"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
	masterdata "source-base-go/usecase/master_data"

	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

func getStatusByType(ctx *gin.Context, masterDataService masterdata.UseCase) {
	
	typeStatus := ctx.Query("type")
	if typeStatus == "" {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}
	listStatus, err := masterDataService.FindStatusByType(typeStatus)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, err)
		return
	}

	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: i18n.MustGetMessage(config.SUCCESS),
		Results: convertListStatusToPresenter(listStatus),
	}
	ctx.JSON(http.StatusOK, response)
}