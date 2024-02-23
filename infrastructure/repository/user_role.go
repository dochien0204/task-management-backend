package repository

import (
	"log"
	"source-base-go/entity"

	"gorm.io/gorm"
)

type UserRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) *UserRoleRepository {
	return &UserRoleRepository{
		db: db,
	}
}

func (r UserRoleRepository) WithTrx(trxHandle *gorm.DB) UserRoleRepository {
	if trxHandle == nil {
		log.Println("Transaction Database not found")
		return r
	}

	r.db = trxHandle
	return r
}

func (r UserRoleRepository) AddRoleForUser(userId, roleId int) error {
	userRole := &entity.UserRole{}
	userRole.RoleId = roleId
	userRole.UserId = userId
	err := r.db.Create(&userRole).Error
	if err != nil {
		return err
	}

	return nil
}
