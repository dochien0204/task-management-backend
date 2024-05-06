package task

import (
	payload "source-base-go/api/payload/task"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository"

	"gorm.io/gorm"
)

type TaskRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.TaskRepository

	Create(data *entity.Task) error
	GetListTaskOfProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.Task, error)
}

type StatusRepository interface {
	GetStatusByCodeAndType(typeStatus string, code string) (*entity.Status, error)
	FindByType(typeStatus string) ([]*entity.Status, error)
}

type UseCase interface {
	WithTrx(trxHandle *gorm.DB) Service

	CreateTask(userId int, payload payload.TaskPayload) error
	GetListTaskOfProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.Task, []*entity.Status, error)
}