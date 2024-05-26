package handler

import (
	"sort"
	projectPayload "source-base-go/api/payload/project"
	presenter "source-base-go/api/presenter/project"
	projectPresenter "source-base-go/api/presenter/project"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"
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

		for _, role := range member.User.Role {
			if role.Type == string(define.PROJECT) {
				userPresenter.User.Role = &presenter.Role{
					Id: role.Id,
					Code: role.Code,
					Name: role.Name,
					Type: role.Type,
				}

				break
			}
		}
		listPresenter = append(listPresenter, userPresenter)
	}

	return listPresenter
}

func convertListActivityProjectByDateToPresenter(listActivity []*entity.Activity) []*projectPresenter.ListActivityProjectByDate {
	listPresenter := []*projectPresenter.ListActivityProjectByDate{}
	mapDate := map[string][]*projectPresenter.Activity{}
	for _, activity := range listActivity {
		//_, isExists := mapDate[task.DueDate.Format(config.DATE_LAYOUT)]
		activityPresenter := &projectPresenter.Activity{
			Id: activity.Id,
			TaskId: activity.TaskId,
			User: &projectPresenter.UserPresenter{
				Id: activity.User.Id,
				Name: activity.User.Name,
				Username: activity.User.Username,
				PhoneNumber: activity.User.PhoneNumber,
				Email: activity.User.Email,
				Avatar: activity.User.Avatar,
			},
			Description: activity.Description,
			CreatedAt: activity.CreatedAt.Format(config.LAYOUT),
			UpdatedAt: activity.UpdatedAt.Format(config.LAYOUT),
		}

		mapDate[activity.CreatedAt.Format(config.DATE_LAYOUT)] = append(mapDate[activity.CreatedAt.Format(config.DATE_LAYOUT)], activityPresenter)
	}

	for k, v := range mapDate {
		listTaskPresenterDetail := &projectPresenter.ListActivityProjectByDate{
			Date: k,
			ListActivity: v,
		}

		listPresenter = append(listPresenter, listTaskPresenterDetail)
	}

	sort.Slice(listPresenter, func(i, j int) bool {
		return listPresenter[i].Date > listPresenter[j].Date
	})
	return listPresenter
}

func convertUserProjectOverviewToPresenter(userOpenTask *entity.UserTaskCount, userClosedTask *entity.UserTaskCount, userProjectRole *entity.UserProjectRole) *presenter.UserProjectOverview {
	presenter := &projectPresenter.UserProjectOverview{
		User: &projectPresenter.UserPresenter{
			Id: userOpenTask.User.Id,
			Name: userOpenTask.User.Name,
			Username: userOpenTask.User.Username,
			PhoneNumber: userOpenTask.User.PhoneNumber,
			Email: userOpenTask.User.Email,
			Avatar: userOpenTask.User.Avatar,
			Role: &projectPresenter.Role {
				Id: userProjectRole.Role.Id,
				Name: userProjectRole.Role.Name,
				Type: userProjectRole.Role.Type,
				Code: userProjectRole.Role.Code,
			},
		},
		TaskOpenCount: userOpenTask.TaskCount,
		TaskCloseCount: userClosedTask.TaskCount,
	}

	return presenter
}

func convertListProjectEntityAdminToPresenter(listData []*entity.Project) []*projectPresenter.ProjectAdminPresenter {
	listProjectPresenter := []*projectPresenter.ProjectAdminPresenter{}
	for _, item := range listData {
		projectPresenter := &projectPresenter.ProjectAdminPresenter{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			Image:       item.Image,
			CreatedAt:   item.CreatedAt.Format(config.LAYOUT),
			UpdatedAt:   item.UpdatedAt.Format(config.LAYOUT),
			Status: &projectPresenter.Status{
				Id: item.Status.Id,
				Code: item.Status.Code,
				Name: item.Status.Name,
				Type: item.Status.Type,
				Description: item.Status.Description,
			},
			User: &projectPresenter.UserPresenter{
				Id: item.User.Id,
				Username: item.User.Username,
				Name: item.User.Name,
				PhoneNumber: item.User.PhoneNumber,
				Email: item.User.Email,
				Avatar: item.User.Avatar,
			},
		}

		listProjectPresenter = append(listProjectPresenter, projectPresenter)
	}

	return listProjectPresenter
}

func converProjectChartViewToPresenter(mapProjectTotalTask map[entity.Project]int, mapProjectDoneTask map[entity.Project]int, mapProjectMemberCount map[entity.Project]int) []*presenter.ProjectChartOverview {
	listPresenter := []*presenter.ProjectChartOverview{}
	for project := range mapProjectTotalTask {
		presenter := &projectPresenter.ProjectChartOverview{
			Project: &projectPresenter.ProjectPresenter{
				Id: project.Id,
				Name: project.Name,
				Description: project.Description,
				Image: project.Image,
				CreatedAt: project.CreatedAt.Format(config.LAYOUT),
				UpdatedAt: project.UpdatedAt.Format(config.LAYOUT),
			},
			TotalTask: mapProjectTotalTask[project],
			DoneTask: mapProjectDoneTask[project],
			MemberCount: mapProjectMemberCount[project],
		}

		listPresenter = append(listPresenter, presenter)
	}

	sort.Slice(listPresenter, func(i, j int) bool {
		return listPresenter[i].Project.Name < listPresenter[j].Project.Name
	})

	return listPresenter
}