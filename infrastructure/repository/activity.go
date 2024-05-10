package repository

import (
	"log"
	"source-base-go/entity"

	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{
		db: db,
	}
}

func (r ActivityRepository) WithTrx(trxHandle *gorm.DB) ActivityRepository {
	if trxHandle == nil {
		log.Print("Transaction DB not found")
		return r
	}

	r.db = trxHandle
	return r
}

func (r ActivityRepository) CreateActivity(data *entity.Activity) error {
	err := r.db.Model(&entity.Activity{}).
		Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}