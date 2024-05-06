package handler

import (
	presenter "source-base-go/api/presenter/task"
	"source-base-go/config"
	"source-base-go/entity"
)

func convertListTaskToPresenter(listTask []*entity.Task, listStatus []*entity.Status) []*presenter.ListTaskPresenter {
	listPresenter := []*presenter.ListTaskPresenter{}
	mapStatusPresenter := map[int]*presenter.ListTaskPresenter{}
	for _, status := range listStatus {
		presenter := &presenter.ListTaskPresenter{
			Status: &presenter.Status {
				Id: status.Id,
				Name: status.Name,
				Code: status.Code,
				Type: status.Type,
			},
		}
		mapStatusPresenter[status.Id] = presenter
		listPresenter = append(listPresenter, presenter)
	}

	for _, task := range listTask {
		taskPresenter := &presenter.Task{
			Id: task.Id,
			Name: task.Name,
			Description: task.Description,
			Category: &presenter.Category{
				Id: task.Category.Id,
				Name: task.Category.Name,
				Code: task.Category.Code,
				Type: task.Category.Type,
			},
		}
		mapStatusPresenter[task.StatusId].ListTask = append(mapStatusPresenter[task.StatusId].ListTask, taskPresenter)
	}

	return listPresenter
}

func convertTaskDetailEntityToPresenter(task *entity.Task) presenter.TaskDetail {
	return presenter.TaskDetail{
		Id: task.Id,
		Name: task.Name,
		Description: task.Description,
		Category: &presenter.Category{
			Id: task.Category.Id,
			Name: task.Category.Name,
			Code: task.Category.Code,
			Type: task.Category.Type,
		},
		User: &presenter.User{
			Id: task.User.Id,
			Username: task.User.Username,
			Name: task.User.Name,
		},
		Assignee: &presenter.User{
			Id: task.User.Id,
			Username: task.User.Username,
			Name: task.User.Name,
		},
		Reviewer: &presenter.User{
			Id: task.User.Id,
			Username: task.User.Username,
			Name: task.User.Name,
		},
		Status: &presenter.Status{
			Id: task.Status.Id,
			Name: task.Status.Name,
			Code: task.Status.Code,
			Type: task.Status.Type,
		},
		StartDate: task.StartDate.Format(config.LAYOUT),
		DueDate: task.DueDate.Format(config.LAYOUT),
		CreatedAt: task.CreatedAt.Format(config.LAYOUT),
		UpdatedAt: task.UpdatedAt.Format(config.LAYOUT),
	}
}