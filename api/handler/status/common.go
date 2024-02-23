package handler

import (
	statusPresenter "source-base-go/api/presenter/status"
	"source-base-go/entity"
)

func ConvertStatusEntityToPresenter(status *entity.Status) *statusPresenter.StatusPresenter {
	return &statusPresenter.StatusPresenter{
		Id:          status.Id,
		Code:        status.Code,
		Name:        status.Name,
		Type:        status.Type,
		Description: status.Description,
	}
}
