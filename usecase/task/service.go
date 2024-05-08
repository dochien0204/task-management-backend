package task

import (
	"fmt"
	payload "source-base-go/api/payload/task"
	taskPayload "source-base-go/api/payload/task"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	taskRepo TaskRepository
	statusRepo StatusRepository
	taskDocumentRepo TaskDocumentRepository
}

func NewService(taskRepo TaskRepository, statusRepo StatusRepository, taskDocumentRepo TaskDocumentRepository) *Service {
	return &Service{
		taskRepo: taskRepo,
		statusRepo: statusRepo,
		taskDocumentRepo: taskDocumentRepo,
	}
}

func (s Service) WithTrx(trxHandle *gorm.DB) Service {
	s.taskRepo = s.taskRepo.WithTrx(trxHandle)
	s.taskDocumentRepo = s.taskDocumentRepo.WithTrx(trxHandle)

	return s
}

func (s Service) CreateTask(userId int, payload taskPayload.TaskPayload) error {
	startDate, err := time.Parse(config.LAYOUT, payload.StartDate)
	dueDate, err := time.Parse(config.LAYOUT, payload.DueDate)
	if err != nil {
		return err
	}
	//Status doing
	status, err := s.statusRepo.GetStatusByCodeAndType(define.TASK_CODE, define.TASK_BACKLOG_STATUS)
	if err != nil {
		return err
	}

	data := &entity.Task {
		Name: payload.Name,
		Description: payload.Description,
		AssigneeId: payload.AssigneeId,
		ReviewerId: payload.ReviewerId,
		CategoryId: payload.CategoryId,
		ProjectId: payload.ProjectId,
		UserId: userId,
		StartDate: startDate,
		DueDate: dueDate,
		StatusId: status.Id,
	}

	err = s.taskRepo.Create(data)
	if err != nil {
		return err
	}

	//Create task document
	listTaskDocument := []*entity.TaskDocument{}
	for _, documentPayload := range payload.Documents {
		path := fmt.Sprintf("task/%d/%v", data.Id, documentPayload.File)
		fileName := fmt.Sprintf("https://%v.s3.%v.amazon.com/%v", config.S3_BUCKET_NAME, config.REGION, path)
		document := &entity.TaskDocument{
			File: fileName,
			Name: documentPayload.Name,
			TaskId: data.Id,
		}

		listTaskDocument = append(listTaskDocument, document)
	}

	err = s.taskDocumentRepo.CreateDocumentsForTask(listTaskDocument)
	if err != nil {
		return err
	}
	
	return nil
}

func (s Service) GetListTaskOfProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.Task, []*entity.Status, error) {
	listStatusTask, err := s.statusRepo.FindByType(define.TASK_CODE)
	if err != nil {
		return nil, nil, err
	}

	listTask, err := s.taskRepo.GetListTaskOfProject(projectId, page, size, sortBy, sortType)
	if err != nil {
		return nil, nil, err
	}

	return listTask, listStatusTask, nil 
}

func (s Service) GetTaskDetail(taskId int) (*entity.Task, error) {
	return s.taskRepo.GetTaskDetail(taskId)
}

func (s Service) UpdateTask(data payload.TaskUpdatePayload) error {
	mapTask := map[string]interface{}{}
	mapTask["name"] = data.Name
	mapTask["description"] = data.Description
	mapTask["assignee_id"] = data.AssigneeId
	mapTask["reviewer_id"] = data.ReviewerId
	mapTask["category_id"] = data.CategoryId
	mapTask["status_id"] = data.StatusId
	startDate, _ := time.Parse(config.LAYOUT, data.StartDate)
	dueDate, _ := time.Parse(config.LAYOUT, data.DueDate)
	mapTask["start_date"]= startDate
	mapTask["due_date"] = dueDate

	return s.taskRepo.UpdateTask(data.Id, mapTask)
}

func (s Service) UpdateTaskStatus(taskId, statusId int) error {
	return s.taskRepo.UpdateStatusTask(taskId, statusId)
}