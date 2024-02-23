package role

import "source-base-go/entity"

type RoleRepository interface {
	GetAllRole() ([]*entity.Role, error)
	FindByCode(code string, typeRole string) (*entity.Role, error)
}

type UseCase interface {
	GetAllRole() ([]*entity.Role, error)
	FindByCode(code string, typeRole string) (*entity.Role, error)
}
