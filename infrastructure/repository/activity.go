package repository

import (
	"fmt"
	"log"
	"source-base-go/entity"
	"time"

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

func (r ActivityRepository) GetListActivityByDate(projectId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Activity, error) {
	listActivity := []*entity.Activity{}
	chain := r.db.Model(&entity.Activity{}).
		Joins("join task t on t.id = activity.task_id AND t.deleted_at is NULL").
		Where("t.project_id = ?", projectId)

	if !fromDate.IsZero() {
		chain = chain.Where(fmt.Sprintf(`(activity.created_at) + interval '%v hour' >= ?`, timeOffset), fromDate)
	}

	if !toDate.IsZero() {
		chain = chain.Where(fmt.Sprintf(`(activity.created_at) + interval '%v hour' <= ?`, timeOffset), toDate)
	}

	err := chain.
		Preload("User").
		Order("created_at DESC").
		Find(&listActivity).Error

	if err != nil {
		return nil, err
	}

	return listActivity, nil
}

func (r ActivityRepository) GetListActivityByDateOfUser(projectId, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Activity, error) {
	listActivity := []*entity.Activity{}
	chain := r.db.Model(&entity.Activity{}).
		Joins("join task t on t.id = activity.task_id AND t.deleted_at is NULL").
		Where("t.project_id = ?", projectId).
		Where("activity.user_id = ?", userId)

	if !fromDate.IsZero() {
		chain = chain.Where(fmt.Sprintf(`(activity.created_at) + interval '%v hour' >= ?`, timeOffset), fromDate)
	}

	if !toDate.IsZero() {
		chain = chain.Where(fmt.Sprintf(`(activity.created_at) + interval '%v hour' <= ?`, timeOffset), toDate)
	}

	err := chain.
		Preload("User").
		Order("created_at desc").
		Find(&listActivity).Error

	if err != nil {
		return nil, err
	}

	return listActivity, nil
}