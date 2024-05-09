package project

import (
	"source-base-go/entity"
	"source-base-go/infrastructure/repository"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.ProjectRepository

	FindProjectByName(name string) (*entity.Project, error)
	CreateProject(data *entity.Project) error
	GetListProjectOfUser(userId, page, size int, sortType, sortBy string) ([]*entity.Project, error)
	GetListMemberByProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.UserTaskCount, error)
	CountListMemberByProject(projectId int) (int, error)
}

type UserProjectRoleRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.UserProjectRoleRepository

	CreateUserProjectRole(data *entity.UserProjectRole) error
	CreateListUserProjectRole(listUserProjectRole []*entity.UserProjectRole) error
	GetProjectOwner(projectId, roleId int) (*entity.UserProjectRole, error)
	GetProjectDetailWithOwner(projectId, roleId int) (*entity.UserProjectRole, error)
}

type RoleRepository interface {
	FindByCode(code string, typeRole string) (*entity.Role, error)
}

type UseCase interface {
	WithTrx(trxHandle *gorm.DB) Service

	CreateProject(userId int, project *entity.Project) error
	GetListProjectOfUser(userId, page, size int, sortType, sortBy string) ([]*entity.Project, error)
	GetProjectDetail(projectId int) (*entity.UserProjectRole, error)
	GetListMemberByProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.UserTaskCount, int, error)
}
