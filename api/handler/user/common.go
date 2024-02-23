package handler

import (
	statusHandler "source-base-go/api/handler/status"
	userPresenter "source-base-go/api/presenter/user"
	"source-base-go/entity"
)

func convertUserEntityToPresenter(user *entity.User) *userPresenter.UserProfilePresenter {
	return &userPresenter.UserProfilePresenter{
		Id:       user.Id,
		Username: user.Username,
		Status:   statusHandler.ConvertStatusEntityToPresenter(user.Status),
	}
}
