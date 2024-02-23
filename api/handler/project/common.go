package handler

import (
	projectPayload "source-base-go/api/payload/project"
	"source-base-go/entity"
)

func convertProjectPayloadToEntity(payload projectPayload.ProjectPayload) *entity.Project {
	return &entity.Project{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
	}
}
