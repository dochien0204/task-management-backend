package handler

import (
	"source-base-go/api/middleware"
	"source-base-go/infrastructure/repository/define"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/user"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, userService user.UseCase, verifier util.Verifier, tx middleware.TxMiddleware) {
	userGroup := app.Group("api/user")
	{
		userGroup.GET("/profile", func(ctx *gin.Context) {
			getUserProfile(ctx, userService)
		})
		userGroup.GET("/list", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			getListUser(ctx, userService)
		})
		userGroup.GET("/avatar/presign-link", func(ctx *gin.Context) {
			getPresignPutURLS3(ctx)
		})
		userGroup.PUT("/update-avatar", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			updateAvatar(ctx, userService)
		})
		userGroup.GET("get-avatar", func(ctx *gin.Context) {
			getAvatarUrl(ctx)
		})
		userGroup.DELETE("/delete", middleware.JWTVerifyMiddleware(verifier), middleware.PermissionMiddleware(define.ADMIN), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			deleteUser(ctx, userService)
		})
		userGroup.PUT("/update", middleware.JWTVerifyMiddleware(verifier), middleware.PermissionMiddleware(define.ADMIN), tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			updateUser(ctx, userService)
		})
	}
}
