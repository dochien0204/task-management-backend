package repository

import (
	"fmt"
	"log"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"

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

func (r DiscussionRepository) GetListDiscussion(taskId, page, size int, sortBy, sortType string) ([]*entity.Discussion, error) {
	listDiscussion := []*entity.Discussion{}
	offset := util.CalculateOffset(page, size)
	if sortType == "" && sortBy == "" {
		sortType = "DESC"
		sortBy = "discussion.created_at"
	}
	if sortType == "" {
		sortType = "DESC"
	}

	switch sortBy {
	case "createdAt":
		sortBy = "discussion.created_at"

	case "updatedAt":
		sortBy = "discussion.updated_at"

	default:
		sortBy = "discussion.created_at"
	}

	err := r.db.Model(&entity.Discussion{}).
		Preload("User").
		Where("task_id = ?", taskId).
		Offset(offset).
		Limit(size).
		Order(fmt.Sprintf("%v %v", sortBy, sortType)).
		Find(&listDiscussion).Error

	if err != nil {
		return nil, err
	}

	return listDiscussion, nil
}

func (r DiscussionRepository) CountListDiscussion(taskId int) (int, error) {
	var count int64
	err := r.db.Model(&entity.Discussion{}).
		Where("task_id = ?", taskId).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return int(count), nil
}