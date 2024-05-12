package project

import (
	"source-base-go/entity"
	"source-base-go/infrastructure/repository"
	"time"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.ProjectRepository

	FindProjectByName(name string) (*entity.Project, error)
	CreateProject(data *entity.Project) error
	GetListProjectOfUser(userId, page, size int, sortType, sortBy string) ([]*entity.Project, error)
	GetListMemberByProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.UserTaskCount, error)
	CountListMemberByProject(projectId int) (int, error)
	CountListTaskOpenUser(projectId, userId, statusId int) (*entity.UserTaskCount, error)
	CountListTaskByStatus(projectId, userId, statusId int) (*entity.UserTaskCount, error)
}

type UserProjectRoleRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.UserProjectRoleRepository

	CreateUserProjectRole(data *entity.UserProjectRole) error
	CreateListUserProjectRole(listUserProjectRole []*entity.UserProjectRole) error
	GetProjectOwner(projectId, roleId int) (*entity.UserProjectRole, error)
	GetProjectDetailWithOwner(projectId, roleId int) (*entity.UserProjectRole, error)
	GetRoleUserInProject(projectId, userId int) (*entity.UserProjectRole, error)
}

type RoleRepository interface {
	FindByCode(code string, typeRole string) (*entity.Role, error)
}

type ActivityRepository interface {
	GetListActivityByDate(projectId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Activity, error)
	GetListActivityByDateOfUser(projectId, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Activity, error)
}

type StatusRepository interface {
	GetStatusByCodeAndType(typeStatus string, code string) (*entity.Status, error)
}

type UseCase interface {
	WithTrx(trxHandle *gorm.DB) Service

	CreateProject(userId int, project *entity.Project) error
	GetListProjectOfUser(userId, page, size int, sortType, sortBy string) ([]*entity.Project, error)
	GetProjectDetail(projectId int) (*entity.UserProjectRole, error)
	GetListMemberByProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.UserTaskCount, int, error)
	GetListActivityProjectByDate(projectId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Activity, error)
	GetOverviewUserTaskProject(projectId, userId int) (*entity.UserTaskCount, *entity.UserTaskCount, *entity.UserProjectRole, error)
	GetListActivityByDateOfUser(projectId, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Activity, error)
}
