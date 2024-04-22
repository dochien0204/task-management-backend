package handler

import (
	presenter "source-base-go/api/presenter/status"
	"source-base-go/entity"
)

func convertListStatusToPresenter(listData []*entity.Status) []*presenter.StatusPresenter {
	listPresenter := []*presenter.StatusPresenter{}
	for _, data := range listData {
		statusPresenter := &presenter.StatusPresenter{
			Id: data.Id,
			Code: data.Code,
			Name: data.Name,
			Type: data.Type,
			Description: data.Description,
		}

		listPresenter = append(listPresenter, statusPresenter)
	}

	return listPresenter
}