package task

import (
	payload "source-base-go/api/payload/task"
	taskPayload "source-base-go/api/payload/task"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository"
	"time"

	"gorm.io/gorm"
)

type TaskRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.TaskRepository

	Create(data *entity.Task) (error)
	GetListTaskOfProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.Task, error)
	GetTaskDetail(taskId int) (*entity.Task, error)
	UpdateTask(taskId int, mapData map[string]interface{}) error
	UpdateStatusTask(taskId int, statusId int) error
	GetListTaskByDate(projectId int, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Task, error)
	GetListTaskProjectByUserAndStatus(projectId int, assigneeId, statusId int, page, size int, sortType, sortBy string) ([]*entity.Task, error)
	CountListTaskProjectByUserAndStatus(projectId int, assigneeId, statusId int) (int, error)
	GetListTaskProjectByUser(projectId int, assigneeId, page, size int, sortType, sortBy string) ([]*entity.Task, error)
	CountListTaskProjectByUser(projectId int, assigneeId int) (int, error)
}

type StatusRepository interface {
	GetStatusByCodeAndType(typeStatus string, code string) (*entity.Status, error)
	FindByType(typeStatus string) ([]*entity.Status, error)
}

type TaskDocumentRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.TaskDocumentRepository

	CreateDocumentsForTask(listData []*entity.TaskDocument) error
	DeleteAllTaskDocumentOfTask(taskId int) error
}

type DiscussionRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.DiscussionRepository

	Create(data *entity.Discussion) error
	GetListDiscussion(taskId, page, size int, sortBy, sortType string) ([]*entity.Discussion, error)
	CountListDiscussion(taskId int) (int, error)
}

type ActivityRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.ActivityRepository

	CreateActivity(data *entity.Activity) error
}

type UserRepository interface {
	FindById(id int ) (*entity.User, error)
}

type ProjectRepository interface {
	FindById(id int ) (*entity.Project, error)
}

type UserProjectRoleRepository interface {
	FindAllUserOfProject(projectId int) ([]*entity.User, error)
}

type EmailRepository interface {
	SendMailForUsers(body string, to []string, subject string) error
}

type UseCase interface {
	WithTrx(trxHandle *gorm.DB) Service

	CreateTask(userId int, payload taskPayload.TaskPayload) (int, error)
	GetListTaskOfProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.Task, []*entity.Status, error)
	GetTaskDetail(taskId int) (*entity.Task, error)
	UpdateTask(payload payload.TaskUpdatePayload) error
	UpdateTaskStatus(taskId, statusId int) error
	GetListTaskByDate(projectId int, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Task, error)
	CreateDiscussionTask(userId, taskId int, comment string) error
	GetListDiscussionOfTask(taskId int, page, size int, sortBy, sortType string) ([]*entity.Discussion, int, error)
	GetListTaskProjectByUserAndStatus(projectId int, assigneeId, statusId int, page, size int, sortType, sortBy string) ([]*entity.Task, int, error)
}