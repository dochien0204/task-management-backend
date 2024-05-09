package repository

import (
	"log"
	"source-base-go/entity"

	"gorm.io/gorm"
)

type TaskDocumentRepository struct {
	db *gorm.DB
}

func NewTaskDocumentRepository(db *gorm.DB) *TaskDocumentRepository {
	return &TaskDocumentRepository{
		db: db,
	}
}

func (r TaskDocumentRepository) WithTrx(trxHandle *gorm.DB) TaskDocumentRepository {
	if trxHandle == nil {
		log.Print("Transaction DB not found")
		return r
	}

	r.db = trxHandle
	return r
}

func (r TaskDocumentRepository) CreateDocumentsForTask(listData []*entity.TaskDocument) error {
	err := r.db.Model(&entity.TaskDocument{}).
		Create(&listData).Error
	
	if err != nil {
		return err
	}

	return nil
}

func (r TaskDocumentRepository) DeleteAllTaskDocumentOfTask(taskId int) error {
	err := r.db.
		Where("task_id = ?", taskId).
		Delete(&entity.TaskDocument{}).Error

	if err != nil {
		return err
	}

	return nil
}