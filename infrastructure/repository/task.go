package repository

import (
	"log"
	"source-base-go/entity"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r TaskRepository) WithTrx(trxHandle *gorm.DB) TaskRepository {
	if trxHandle == nil {
		log.Print("Transaction DB not found")
		return r
	}

	r.db = trxHandle
	return r
}

func (r TaskRepository) Create(data *entity.Task) error {
	err := r.db.Model(&entity.Task{}).
		Create(&data).Error
	
	if err != nil {
		return err
	}

	return nil
}