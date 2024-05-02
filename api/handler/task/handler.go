package handler

import (
	"source-base-go/api/middleware"
	"source-base-go/infrastructure/repository/define"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/task"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, taskService task.UseCase, verifier util.Verifier, tx middleware.TxMiddleware) {
	taskGroup := app.Group("api/task")
	{
		taskGroup.POST("/create", middleware.JWTVerifyMiddleware(verifier), middleware.PermissionMiddleware(define.ADMIN), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			createTask(ctx, taskService)
		})
	}
}
