package handler

import (
	"fmt"
	"net/http"
	payload "source-base-go/api/payload/user"
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

func updateAvatar(ctx *gin.Context, userService user.UseCase) {
	var payload payload.UpdateAvatar

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandleException(ctx, http.StatusOK, entity.ErrBadRequest)
		return
	}

	token, err := util.GetToken(ctx)
	if err != nil {
		util.HandleException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
		return
	}

	claims, err := util.ParseAccessToken(token)
	if err != nil {
		util.HandleException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
		return
	}

	err = userService.UpdateAvatar(claims.UserId, payload.Avatar)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	response := presenter.BasicResponse{
		Status: fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
	}

	ctx.JSON(http.StatusOK, response)
}

func getAvatarUrl(ctx *gin.Context) {
	token, err := util.GetToken(ctx)
	if err != nil {
		util.HandleException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
		return
	}

	_, err = util.ParseAccessToken(token)
	if err != nil {
		util.HandleException(ctx, http.StatusUnauthorized, entity.ErrUnauthorized)
		return
	}

	avatar := ctx.Query("avatar")
	userId := ctx.Query("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	urlView, err := util.GeneratePresignViewAvatarURLS3(userIdInt, avatar)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	response := presenter.BasicResponse{
		Status: fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
		Results: map[string]interface{}{
			"url": urlView,
		},
	}

	ctx.JSON(http.StatusOK, response)
}