package repository

import (
	"source-base-go/entity"

	"gorm.io/gorm"
)

type StatusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) *StatusRepository {
	return &StatusRepository{
		db: db,
	}
}

func (r StatusRepository) FindByType(typeStatus string) ([]*entity.Status, error) {
	listStatus := []*entity.Status{}
	err := r.db.Model(&entity.Status{}).
		Where("type = ?", typeStatus).
		Find(&listStatus).Error
	
	if err != nil {
		return nil, err
	}

	return listStatus, nil
}