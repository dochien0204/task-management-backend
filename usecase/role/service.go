package role

import "source-base-go/entity"

type Service struct {
	roleRepo RoleRepository
}

func NewService(roleRepo RoleRepository) *Service {
	return &Service{
		roleRepo: roleRepo,
	}
}

func (s Service) GetAllRole() ([]*entity.Role, error) {
	return s.roleRepo.GetAllRole()
}

func (s Service) FindByCode(code string, typeRole string) (*entity.Role, error) {
	return s.roleRepo.FindByCode(code, typeRole)
}

func (s Service) FindAllRoleByType(typeRole string) ([]*entity.Role, error) {
	return s.roleRepo.FindAllRoleByType(typeRole)
}
