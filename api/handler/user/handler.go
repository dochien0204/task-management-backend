package handler

import (
	"source-base-go/api/middleware"
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
	}
}
