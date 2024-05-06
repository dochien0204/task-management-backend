package handler

import (
	presenter "source-base-go/api/presenter/task"
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