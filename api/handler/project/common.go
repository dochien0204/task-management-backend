package handler

import (
	projectPayload "source-base-go/api/payload/project"
	projectPresenter "source-base-go/api/presenter/project"
	"source-base-go/config"
	"source-base-go/entity"
)

func convertProjectPayloadToEntity(payload projectPayload.ProjectPayload) *entity.Project {
	return &entity.Project{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
	}
}

func convertListProjectEntityToPresenter(listData []*entity.Project) []*projectPresenter.ProjectPresenter {
	listProjectPresenter := []*projectPresenter.ProjectPresenter{}
	for _, item := range listData {
		projectPresenter := &projectPresenter.ProjectPresenter{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			Image:       item.Image,
			CreatedAt:   item.CreatedAt.Format(config.LAYOUT),
			UpdatedAt:   item.UpdatedAt.Format(config.LAYOUT),
		}

		listProjectPresenter = append(listProjectPresenter, projectPresenter)
	}

	return listProjectPresenter
}
