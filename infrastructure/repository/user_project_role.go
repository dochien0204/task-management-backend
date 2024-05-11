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

func (r UserProjectRoleRepository) GetProjectDetailWithOwner(projectId, roleId int) (*entity.UserProjectRole, error) {
	userProjectRole := &entity.UserProjectRole{}
	err := r.db.Model(&entity.UserProjectRole{}).
		Where("project_id = ?", projectId).
		Where("role_id = ?", roleId).
		Preload("Project").
		Preload("User").
		First(&userProjectRole).Error
	if err != nil {
		return nil, err
	}

	return userProjectRole, nil
}

func (r UserProjectRoleRepository) FindAllUserOfProject(projectId int) ([]*entity.User, error) {
	listUser := []*entity.User{}
	err := r.db.Model(&entity.UserProjectRole{}).
		Select("u.*").
		Joins(`join "user" u on u.id = user_project_role.user_id`).
		Where("user_project_role.project_id = ?", projectId).
		Find(&listUser).Error

	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (r UserProjectRoleRepository) GetRoleUserInProject(projectId, userId int) (*entity.UserProjectRole, error) {
	result := &entity.UserProjectRole{}
	err := r.db.Model(&entity.UserProjectRole{}).
		Preload("Role").
		Where("user_id = ?", userId).
		Where("project_id = ?", projectId).
		First(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
