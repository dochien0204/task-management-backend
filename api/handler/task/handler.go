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

		taskGroup.GET("/list", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			getListTaskOfProject(ctx, taskService)
		})

		taskGroup.GET("/detail", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			getTaskDetail(ctx, taskService)
		})

		taskGroup.PUT("/update", middleware.JWTVerifyMiddleware(verifier), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			updateTask(ctx, taskService)
		})
		taskGroup.PUT("/update-status", middleware.JWTVerifyMiddleware(verifier), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			updateTaskStatus(ctx, taskService)
		})

		taskGroup.GET("/list-task-by-date", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			getListTaskByDate(ctx, taskService)
		})

		taskGroup.POST("/discussion", tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			createDiscussion(ctx, taskService)
		})

		taskGroup.GET("/discussion/list", tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			getListDiscussionByTask(ctx, taskService)
		})
	}
}
