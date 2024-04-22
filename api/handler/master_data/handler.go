package handler

import (
	"source-base-go/infrastructure/repository/util"
	masterdata "source-base-go/usecase/master_data"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, masterDataService masterdata.UseCase, verifier util.Verifier) {
	projectGroup := app.Group("api/master-data")
	{
		projectGroup.GET("/status", func(ctx *gin.Context) {
			getStatusByType(ctx, masterDataService)
		})
	}
}