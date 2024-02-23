package handler

import (
	"fmt"
	"net/http"
	authPayload "source-base-go/api/payload/auth"
	"source-base-go/api/presenter"
	authPresenter "source-base-go/api/presenter/auth"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/user"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func login(ctx *gin.Context, authService user.UseCase) {
	var payload authPayload.Login
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	//Login
	tokenPair, user, err := authService.Login(payload.Username, payload.Password)
	if err != nil {
		switch err {
		case entity.ErrUsernameNotExists:
			util.HandleException(ctx, http.StatusBadRequest, err)
			return
		case entity.ErrInvalidPassword:
			util.HandleException(ctx, http.StatusBadRequest, err)
			return
		default:
			util.HandleException(ctx, http.StatusBadRequest, err)
			return
		}
	}

	result := &authPresenter.AuthResult{
		AccessToken:  tokenPair.Token,
		RefreshToken: tokenPair.RefreshToken,
		UserId:       user.Id,
		Username:     user.Username,
	}

	//Response in JSON
	response := &authPresenter.AuthResp{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
		Results: result,
	}

	ctx.JSON(http.StatusOK, response)
}

func register(ctx *gin.Context, authService user.UseCase) {
	var payload authPayload.Register
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	//Get transaction
	txHandle := ctx.MustGet("db_trx").(*gorm.DB)

	err = authService.WithTrx(txHandle).Register(convertRegisterPayloadToEntity(&payload))
	if err != nil {
		switch err {
		case entity.ErrBadRequest:
			util.HandleException(ctx, http.StatusBadRequest, err)
			return
		case entity.ErrAccountAlreadyExists:
			util.HandleException(ctx, http.StatusBadRequest, err)
			return
		default:
			util.HandleException(ctx, http.StatusInternalServerError, entity.ErrInternalServerError)
			return
		}
	}

	//Response in JSON
	response := presenter.BasicResponse{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
	}

	ctx.JSON(http.StatusOK, response)
}
