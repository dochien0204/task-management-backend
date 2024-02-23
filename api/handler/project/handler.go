package handler

import (
	"source-base-go/api/middleware"
	"source-base-go/infrastructure/repository/define"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/project"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, projectService project.UseCase, verifier util.Verifier, tx middleware.TxMiddleware) {
	projectGroup := app.Group("api/project")
	{
		projectGroup.POST("/create", middleware.JWTVerifyMiddleware(verifier), middleware.PermissionMiddleware(define.ADMIN), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			createProject(ctx, projectService)
		})

		projectGroup.GET("/list-project", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			getListProjectOfUser(ctx, projectService)
		})

		projectGroup.POST("/add-member-to-project", middleware.JWTVerifyMiddleware(verifier), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			addListMemberToProject(ctx, projectService)
		})
	}
}