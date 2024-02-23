package handler

import (
	"source-base-go/api/middleware"
	"source-base-go/usecase/user"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, authService user.UseCase, tx middleware.TxMiddleware) {
	authGroup := app.Group("/api/auth")
	{
		authGroup.POST("/login", func(ctx *gin.Context) {
			login(ctx, authService)
		})
		authGroup.POST("/register", tx.DBTransactionMiddleware(), func(ctx *gin.Context) {
			register(ctx, authService)
		})
	}
}
