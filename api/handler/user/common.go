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

func convertListUserToPresenter(listUser []*entity.User) []*userPresenter.UserPresenter {
	listPresenster := []*userPresenter.UserPresenter{}
	for _, user := range listUser {
		userPresenter := &userPresenter.UserPresenter{
			Id: user.Id,
			Username: user.Username,
			Name: user.Name,
			PhoneNumber: user.PhoneNumber,
			Email: user.Email,
			Avatar: user.Avatar,
		}

		listPresenster = append(listPresenster, userPresenter)
	}

	return listPresenster
}
