package repository

import (
	"fmt"
	"log"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"

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

func (r TaskRepository) GetListTaskOfProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.Task, error) {
	offset := util.CalculateOffset(page, size)
	if sortType == "" && sortBy == "" {
		sortType = "DESC"
		sortBy = "project.created_at"
	}
	if sortType == "" {
		sortType = "DESC"
	}

	switch sortBy {
	case "createdAt":
		sortBy = "task.created_at"

	case "updatedAt":
		sortBy = "task.updated_at"

	default:
		sortBy = "task.created_at"
	}

	listTask := []*entity.Task{}
	err := r.db.Model(&entity.Task{}).
		Preload("User").
		Preload("Assignee").
		Preload("Category").
		Where("project_id = ?", projectId).
		Offset(offset).
		Limit(size).
		Order(fmt.Sprintf("%v %v", sortBy, sortType)).
		Find(&listTask).Error
	
	if err != nil {
		return nil, err
	}

	return listTask, nil
}

func (r TaskRepository) GetTaskDetail(taskId int) (*entity.Task, error) {
	task := &entity.Task{}
	err := r.db.Model(&entity.Task{}).
		Preload("User").
		Preload("Assignee").
		Preload("Category").
		Preload("Status").
		Where("id = ?", taskId).
		First(&task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r TaskRepository) UpdateTask(taskId int, mapData map[string]interface{}) error {
	err := r.db.Model(&entity.Task{}).
		Where("id = ?", taskId).
		Updates(mapData).Error
	if err != nil {
		return err
	}

	return nil
} 