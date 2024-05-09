package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDataFromHeader(ctx *gin.Context, headerName string) int {
	listTimeOffset := ctx.Request.Header[headerName]

	timeOffsetStr := ""
	if len(listTimeOffset) > 0 {
		timeOffsetStr = listTimeOffset[0]
	}

	timeOffset, _ := strconv.Atoi(timeOffsetStr)

	return timeOffset
}