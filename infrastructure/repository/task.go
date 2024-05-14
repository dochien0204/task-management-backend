package repository

import (
	"fmt"
	"log"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
	"time"

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

func (r TaskRepository) Create(data *entity.Task) (error) {
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
		Preload("Reviewer").
		Preload("Category").
		Preload("Status").
		Preload("Documents").
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

func (r TaskRepository) UpdateStatusTask(taskId int, statusId int) error {
	err := r.db.Model(&entity.Task{}).
		Where("id = ?", taskId).
		Update("status_id", statusId).Error
	if err != nil {
		return err
	}

	return nil
}


func (r TaskRepository) GetListTaskByDate(projectId int, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Task, error) {
	listTask := []*entity.Task{}
	chain := r.db.Model(&entity.Task{}).
		Where("project_id = ?", projectId).
		Where("assignee_id = ?", userId)

	if !fromDate.IsZero() {
		chain = chain.Where(fmt.Sprintf(`(task.due_date) + interval '%v hour' >= ?`, timeOffset), fromDate)
	}

	if !toDate.IsZero() {
		chain = chain.Where(fmt.Sprintf(`(task.due_date) + interval '%v hour' <= ?`, timeOffset), toDate)
	}

	err := chain.
		Preload("User").
		Preload("Assignee").
		Preload("Reviewer").
		Preload("Category").
		Preload("Status").
		Preload("Documents").
		Order("created_at desc").
		Find(&listTask).Error

	if err != nil {
		return nil, err
	}

	return listTask, nil
}	

func (r TaskRepository) GetListTaskProjectByUserAndStatus(projectId int, assigneeId, statusId int, page, size int, sortType, sortBy string) ([]*entity.Task, error) {
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
		Preload("Status").
		Where("project_id = ?", projectId).
		Where("assignee_id = ?", assigneeId).
		Where("status_id = ?", statusId).
		Offset(offset).
		Limit(size).
		Order(fmt.Sprintf("%v %v", sortBy, sortType)).
		Find(&listTask).Error
	
	if err != nil {
		return nil, err
	}

	return listTask, nil
}

func (r TaskRepository) CountListTaskProjectByUserAndStatus(projectId int, assigneeId, statusId int) (int, error) {
	var count int64
	err := r.db.Model(&entity.Task{}).
		Preload("User").
		Preload("Assignee").
		Preload("Category").
		Preload("Status").
		Where("project_id = ?", projectId).
		Where("assignee_id = ?", assigneeId).
		Where("status_id = ?", statusId).
		Count(&count).Error
	
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r TaskRepository) GetListTaskProjectByUser(projectId int, assigneeId, page, size int, sortType, sortBy string) ([]*entity.Task, error) {
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
		Preload("Status").
		Where("project_id = ?", projectId).
		Where("assignee_id = ?", assigneeId).
		Offset(offset).
		Limit(size).
		Order(fmt.Sprintf("%v %v", sortBy, sortType)).
		Find(&listTask).Error
	
	if err != nil {
		return nil, err
	}

	return listTask, nil
}

func (r TaskRepository) CountListTaskProjectByUser(projectId int, assigneeId int) (int, error) {
	var count int64
	err := r.db.Model(&entity.Task{}).
		Preload("User").
		Preload("Assignee").
		Preload("Category").
		Preload("Status").
		Where("project_id = ?", projectId).
		Where("assignee_id = ?", assigneeId).
		Count(&count).Error
	
	if err != nil {
		return 0, err
	}

	return int(count), nil
}