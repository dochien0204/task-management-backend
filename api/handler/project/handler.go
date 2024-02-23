package handler

import (
	"source-base-go/api/middleware"
	"source-base-go/infrastructure/repository/define"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/project"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, projectService project.UseCase, verifier util.Verifier, tx middleware.TxMiddleware) {
	userGroup := app.Group("api/project")
	{
		userGroup.POST("/create", middleware.JWTVerifyMiddleware(verifier), middleware.PermissionMiddleware(define.ADMIN), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			createProject(ctx, projectService)
		})
	}
}
