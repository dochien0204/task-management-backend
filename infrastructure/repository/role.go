package repository

import (
	"log"
	"source-base-go/entity"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r RoleRepository) WithTrx(trxHandle *gorm.DB) RoleRepository {
	if trxHandle == nil {
		log.Println("Transaction Database not found")
		return r
	}

	r.db = trxHandle
	return r
}

func (r RoleRepository) GetAllRole() ([]*entity.Role, error) {
	var listRole []*entity.Role
	err := r.db.Find(&listRole).Error
	if err != nil {
		return nil, err
	}

	return listRole, nil
}

func (r RoleRepository) FindByCode(code string, typeRole string) (*entity.Role, error) {
	var role *entity.Role
	err := r.db.
		Where("code = ?", code).
		Where("type = ?", typeRole).
		Find(&role).Error
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r RoleRepository) FindAllRolesOfUser(userId int) ([]*entity.Role, error) {
	listRole := []*entity.Role{}
	err := r.db.
		Select("role.*").
		Joins("join user_role ur on ur.role_id = id").
		Where("ur.user_id = ?", userId).
		Find(&listRole).Error
	if err != nil {
		return nil, err
	}

	return listRole, nil
}

func (r RoleRepository) FindAllRoleByType(typeRole string) ([]*entity.Role, error) {
	listRole := []*entity.Role{}
	err := r.db.
		Select("role.*").
		Where("type = ?", typeRole).
		Find(&listRole).Error

	if err != nil {
		return nil, err
	}

	return listRole, nil
}
