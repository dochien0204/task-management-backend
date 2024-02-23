package repository

import (
	"log"
	"source-base-go/entity"

	"gorm.io/gorm"
)

type UserProjectRoleRepository struct {
	db *gorm.DB
}

func NewUserProjectRole(db *gorm.DB) *UserProjectRoleRepository {
	return &UserProjectRoleRepository{
		db: db,
	}
}

func (r UserProjectRoleRepository) WithTrx(trxHandle *gorm.DB) UserProjectRoleRepository {
	if trxHandle == nil {
		log.Print("Transaction DB not found")
		return r
	}

	r.db = trxHandle
	return r
}

func (r UserProjectRoleRepository) CreateUserProjectRole(data *entity.UserProjectRole) error {
	err := r.db.Model(&entity.UserProjectRole{}).Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (r UserProjectRoleRepository) CreateListUserProjectRole(listUserProjectRole []*entity.UserProjectRole) error {
	err := r.db.Model(&entity.UserProjectRole{}).Create(&listUserProjectRole).Error
	if err != nil {
		return err
	}

	return nil
}

func (r UserProjectRoleRepository) GetProjectOwner(projectId, roleId int) (*entity.UserProjectRole, error) {
	userProjectRole := &entity.UserProjectRole{}
	err := r.db.Model(&entity.UserProjectRole{}).
		Where("project_id = ?", projectId).
		Where("role_id = ?", roleId).
		Find(&userProjectRole).Error
	if err != nil {
		return nil, err
	}

	return userProjectRole, nil
}
