package util

import (
	"fmt"
	"source-base-go/api/presenter"

	"github.com/gin-gonic/gin"
)

func HandleException(ctx *gin.Context, statusCode int, err error) {
	errorMessage := &presenter.BasicResponse{
		Status:  fmt.Sprint(statusCode),
		Message: ParseError(err),
	}

	ctx.AbortWithStatusJSON(statusCode, errorMessage)
}
