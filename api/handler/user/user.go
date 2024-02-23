package handler

import (
	"fmt"
	"net/http"
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
