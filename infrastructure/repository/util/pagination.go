package util

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

var allowPageSizes = []int{10, 15, 20, 30, 50}

func NormalizePageSize(pageSize int) int {
	allow := false
	for _, allowPageSize := range allowPageSizes {
		if pageSize == allowPageSize {
			allow = true
			break
		}
	}

	if !allow {
		pageSize = allowPageSizes[0]
	}

	return pageSize
}

func CalculateOffset(page int, pageSize int) int {
	offset := 0
	if page != 0 {
		offset += (page - 1) * pageSize
	}

	return offset
}

func CalculateTotalPages(numTotalResults int, pageSize int) float64 {
	return math.Ceil(float64(numTotalResults) / float64(pageSize))
}

func GetPage(ctx *gin.Context, paramName string) int {
	page := ctx.Query(paramName)

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 0 {
		return 1
	}

	return pageInt
}

func GetPageSize(ctx *gin.Context, paramName string) int {
	pageSize := ctx.Query(paramName)

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 0 {
		return 10
	}

	return pageSizeInt
}
