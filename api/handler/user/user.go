package handler

import (
	"fmt"
	"net/http"
	"source-base-go/api/presenter"
	userPresenter "source-base-go/api/presenter/user"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/user"
	"strconv"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

// @Summary Get user's profile
// @Schemes
// @Description Get user's profile
// @Tags user
// @Success 200 {object} userPresenter.UserProfileResponse
// @Router /user/profile [get]
func getUserProfile(ctx *gin.Context, userService user.UseCase) {
	//URL param
	userId := ctx.Query("userId")

	//convert userId param to user id int
	userIdConv, err := strconv.Atoi(userId)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	data, err := userService.GetUserProfile(userIdConv)
	if err != nil {
		util.HandleException(ctx, http.StatusInternalServerError, entity.ErrInternalServerError)
		return
	}

	//Response in JSON
	response := &userPresenter.UserProfilePresenterResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
		Result:  convertUserEntityToPresenter(data),
	}

	ctx.JSON(http.StatusOK, response)
}

func getListUser(ctx *gin.Context, userService user.UseCase) {
	page := util.GetPage(ctx, "page")
	pageSize := util.GetPageSize(ctx, "size")
	sortBy := ctx.Query("sortBy")
	sortType := ctx.Query("sortType")

	listUser, count, err := userService.GetListUser(page, pageSize, sortType, sortBy)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	response := presenter.PaginationResponse {
		Status: fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
		Results: convertListUserToPresenter(listUser),
		Pagination: presenter.Pagination {
			Count: count,
			NumPages: int(util.CalculateTotalPages(count, pageSize)),
			DisplayRecord: len(listUser),
			Page: page,
		},
	}

	ctx.JSON(http.StatusOK, response)
}

func getPresignPutURLS3(ctx *gin.Context) {
	keyName := ctx.Query("keyName")

	presginUrl, err := util.GeneratePresignUploadS3(keyName)

	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, err)
		return
	}

	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
		Results: presginUrl,
	}
	ctx.JSON(http.StatusOK, response)
}