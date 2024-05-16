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

		projectGroup.GET("/detail", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			getProjectDetail(ctx, projectService)
		})

		projectGroup.GET("/task/list-member", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			getListMemberTaskCount(ctx, projectService)
		})

		projectGroup.GET("/activity/list", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			getListActivityProjectByDate(ctx, projectService)
		})

		projectGroup.GET("/member/overview", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			getUserOverviewTaskProject(ctx, projectService)
		})

		projectGroup.GET("/activity/user", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			getListActivityProjectByUser(ctx, projectService)
		})

		projectGroup.PUT("/update", middleware.JWTVerifyMiddleware(verifier), middleware.PermissionMiddleware(define.ADMIN), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			updateProject(ctx, projectService)
		})

		projectGroup.GET("/list/admin", middleware.JWTVerifyMiddleware(verifier), middleware.PermissionMiddleware(define.ADMIN), func(ctx *gin.Context) {
			getAllProject(ctx, projectService)
		})

		projectGroup.DELETE("/delete", middleware.JWTVerifyMiddleware(verifier), middleware.PermissionMiddleware(define.ADMIN), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			deleteProject(ctx, projectService)
		})
		projectGroup.POST("/add-member", middleware.JWTVerifyMiddleware(verifier), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			addListMemberWithRoleToProject(ctx, projectService)
		})
	}
}
