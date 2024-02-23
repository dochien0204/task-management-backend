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
}

type UserProjectRoleRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.UserProjectRoleRepository

	CreateUserProjectRole(data *entity.UserProjectRole) error
}

type RoleRepository interface {
	FindByCode(code string, typeRole string) (*entity.Role, error)
}

type UseCase interface {
	WithTrx(trxHandle *gorm.DB) Service

	CreateProject(userId int, project *entity.Project) error
}
