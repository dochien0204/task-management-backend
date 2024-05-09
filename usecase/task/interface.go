package task

import (
	payload "source-base-go/api/payload/task"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository"
	"time"

	"gorm.io/gorm"
)

type TaskRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.TaskRepository

	Create(data *entity.Task) error
	GetListTaskOfProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.Task, error)
	GetTaskDetail(taskId int) (*entity.Task, error)
	UpdateTask(taskId int, mapData map[string]interface{}) error
	UpdateStatusTask(taskId int, statusId int) error
	GetListTaskByDate(projectId int, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Task, error)
}

type StatusRepository interface {
	GetStatusByCodeAndType(typeStatus string, code string) (*entity.Status, error)
	FindByType(typeStatus string) ([]*entity.Status, error)
}

type TaskDocumentRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.TaskDocumentRepository

	CreateDocumentsForTask(listData []*entity.TaskDocument) error
}

type UseCase interface {
	WithTrx(trxHandle *gorm.DB) Service

	CreateTask(userId int, payload payload.TaskPayload) error
	GetListTaskOfProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.Task, []*entity.Status, error)
	GetTaskDetail(taskId int) (*entity.Task, error)
	UpdateTask(payload payload.TaskUpdatePayload) error
	UpdateTaskStatus(taskId, statusId int) error
	GetListTaskByDate(projectId int, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Task, error)
}