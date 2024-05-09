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

func convertProjectDetailToPresenter(userProject *entity.UserProjectRole) projectPresenter.ProjectDetailPresenter {
	return projectPresenter.ProjectDetailPresenter{
		Id:          userProject.Project.Id,
		Name:        userProject.Project.Name,
		Image:       userProject.Project.Image,
		Description: userProject.Project.Description,
		CreatedAt:   userProject.Project.CreatedAt.Format(config.LAYOUT),
		UpdatedAt:   userProject.Project.UpdatedAt.Format(config.LAYOUT),
		Owner: &projectPresenter.User{
			Id:       userProject.User.Id,
			Username: userProject.User.Username,
			Name:     userProject.User.Name,
		},
	}
}

func convertListMemberTaskCountToPresenter(listMember []*entity.UserTaskCount) []*projectPresenter.UserTaskCount {
	listPresenter := []*projectPresenter.UserTaskCount{}
	for _, member := range listMember {
		userPresenter := &projectPresenter.UserTaskCount{
			User: &projectPresenter.UserPresenter{
				Id: member.User.Id,
				Username: member.User.Username,
				Name: member.User.Name,
				PhoneNumber: member.User.PhoneNumber,
				Email: member.User.Email,
				Avatar: member.User.Avatar,
			},
			TaskCount: member.TaskCount,
		}

		listPresenter = append(listPresenter, userPresenter)
	}

	return listPresenter

}