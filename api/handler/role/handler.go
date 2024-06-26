package handler

import (
	"source-base-go/api/middleware"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/role"

	"github.com/gin-gonic/gin"
)

func MakeHandlers(app *gin.Engine, roleService role.UseCase, verifier util.Verifier, tx middleware.TxMiddleware) {
	roleGroup := app.Group("/api/role")
	{
		roleGroup.GET("/list", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			findAllRole(ctx, roleService)
		})
		roleGroup.GET("/find", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			findByCode(ctx, roleService)
		})
		roleGroup.GET("/get-by-type", middleware.JWTVerifyMiddleware(verifier), func(ctx *gin.Context) {
			findByType(ctx, roleService)
		})
	}
}
