package handler

import (
	"fmt"
	"net/http"
	rolePresenter "source-base-go/api/presenter/role"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/role"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

func findAllRole(ctx *gin.Context, roleService role.UseCase) {
	//Get data
	listRole, err := roleService.GetAllRole()
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	//Response in JSON
	response := &rolePresenter.ListRoleResp{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
		Results: convertListRoleEntityToPresenter(listRole),
	}

	ctx.JSON(http.StatusOK, response)
}

func findByCode(ctx *gin.Context, roleService role.UseCase) {
	//Get URL param
	code := ctx.Query("code")
	typeRole := ctx.Query("type")
	//Get data
	role, err := roleService.FindByCode(code, typeRole)
	if err != nil {
		util.HandleException(ctx, http.StatusBadRequest, entity.ErrBadRequest)
		return
	}

	//Response in JSON
	response := &rolePresenter.RoleResp{
		Status:  fmt.Sprint(http.StatusOK),
		Message: ginI18n.MustGetMessage(config.SUCCESS),
		Results: convertRoleEntityToPresenter(role),
	}

	ctx.JSON(http.StatusOK, response)
}
