package task

import (
	taskPayload "source-base-go/api/payload/task"
	presenter "source-base-go/api/presenter/task"
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
	DeleteTask(taskId int) error
	GetListTaskByUser(listProjectId []int, keyword string, page, size int, sortBy, sortType string) ([]*entity.Task, error)
	CountListTaskByUser(listProjectId []int, keyword string) (int, error)
	CountListRestTaskOfUser(userId int) (int, error)
}

type StatusRepository interface {
	GetStatusByCodeAndType(typeStatus string, code string) (*entity.Status, error)
	FindByType(typeStatus string) ([]*entity.Status, error)
	FindById(id int) (*entity.Status, error)
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
	CountProductOfUser(userId, statusId int) (int, error)
}

type UserProjectRoleRepository interface {
	FindAllUserOfProject(projectId int) ([]*entity.User, error)
	GetRoleUserInProject(projectId, userId int) (*entity.UserProjectRole, error)
	GetAllProjectOfUser(userId int) ([]*entity.UserProjectRole, error)
}

type EmailRepository interface {
	SendMailForUsers(body string, to []string, subject string) error
}

type UseCase interface {
	WithTrx(trxHandle *gorm.DB) Service

	CreateTask(userId int, payload taskPayload.TaskPayload) (int, error)
	GetListTaskOfProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.Task, []*entity.Status, error)
	GetTaskDetail(taskId int) (*entity.Task, error)
	UpdateTask(userId int, data taskPayload.TaskUpdatePayload) error
	UpdateTaskStatus(userId, taskId, statusId int) error
	GetListTaskByDate(projectId int, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Task, error)
	CreateDiscussionTask(userId, taskId int, comment string) error
	GetListDiscussionOfTask(taskId int, page, size int, sortBy, sortType string) ([]*entity.Discussion, int, error)
	GetListTaskProjectByUserAndStatus(projectId int, assigneeId, statusId int, page, size int, sortType, sortBy string) ([]*entity.Task, int, error)
	DeleteTask(userId, taskId int) error
	GetListTaskByUser(userId int, keyword string, page, size int, sortBy, sortType string) ([]*entity.Task, int, error)
	GetOverviewCard(userId int) (*presenter.DashboardCardOverview, error)
}