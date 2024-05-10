package repository

import (
	"log"
	"source-base-go/entity"

	"gorm.io/gorm"
)

type DiscussionRepository struct {
	db *gorm.DB
}

func NewDiscussionRepository(db *gorm.DB) *DiscussionRepository {
	return &DiscussionRepository{
		db: db,
	}
}

func (r DiscussionRepository) WithTrx(trxHandle *gorm.DB) DiscussionRepository {
	if trxHandle == nil {
		log.Print("Transaction DB not found")
		return r
	}

	r.db = trxHandle
	return r
}

func (r DiscussionRepository) Create(data *entity.Discussion) error {
	err := r.db.Model(&entity.Discussion{}).
		Create(&data).Error
	
	if err != nil {
		return err
	}

	return nil
}