package handler

import (
	presenter "source-base-go/api/presenter/task"
	"source-base-go/config"
	"source-base-go/entity"
	"time"
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
	taskPresenter := presenter.TaskDetail{
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

	listDocumentPresenter := []*presenter.Document{}
	for _, document := range task.Documents {
		documentPresenter := &presenter.Document {
			Id: document.Id,
			Name: document.Name,
			FileName: document.File,
			TaskId: document.TaskId,
			CreatedAt: document.CreatedAt.Format(config.LAYOUT),
			UpdatedAt: document.UpdatedAt.Format(time.Layout),
		}

		listDocumentPresenter = append(listDocumentPresenter, documentPresenter)
	}

	taskPresenter.Documents = listDocumentPresenter
	return taskPresenter
}

func convertListTaskByDateToPresenter(listTask []*entity.Task) []*presenter.ListTaskByDatePresenter {
	listPresenter := []*presenter.ListTaskByDatePresenter{}
	mapDate := map[string][]*presenter.TaskDetail{}
	for _, task := range listTask {
		//_, isExists := mapDate[task.DueDate.Format(config.DATE_LAYOUT)]
		taskPresenter := &presenter.TaskDetail{
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

		mapDate[task.DueDate.Format(config.DATE_LAYOUT)] = append(mapDate[task.DueDate.Format(config.DATE_LAYOUT)], taskPresenter)
	}

	for k, v := range mapDate {
		listTaskPresenterDetail := &presenter.ListTaskByDatePresenter{
			Date: k,
			ListTask: v,
			Count: len(v),
		}

		listPresenter = append(listPresenter, listTaskPresenterDetail)
	}

	return listPresenter
}