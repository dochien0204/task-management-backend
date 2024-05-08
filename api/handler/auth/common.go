package handler

import (
	authPayload "source-base-go/api/payload/auth"
	"source-base-go/entity"
)

func convertRegisterPayloadToEntity(data *authPayload.Register) *entity.User {
	return &entity.User{
		Username: data.Username,
		Password: data.Password,
		Email: data.Email,
		PhoneNumber: data.PhoneNumber,
		Avatar: data.Avatar,
	}
}
