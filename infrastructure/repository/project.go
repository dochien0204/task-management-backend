package repository

import (
	"fmt"
	"log"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (r ProjectRepository) WithTrx(trxHandle *gorm.DB) ProjectRepository {
	if trxHandle == nil {
		log.Print("Transaction DB not found")
		return r
	}

	r.db = trxHandle
	return r
}

func (r ProjectRepository) FindProjectByName(name string) (*entity.Project, error) {
	project := &entity.Project{}
	err := r.db.Model(&entity.Project{}).
		Where("name = ?", name).
		Find(&project).Error
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (r ProjectRepository) GetListProjectOfUser(userId, statusId, page, size int, sortType, sortBy string) ([]*entity.Project, error) {
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
		sortBy = "project.created_at"

	case "updatedAt":
		sortBy = "project.updated_at"

	default:
		sortBy = "project.created_at"
	}
	listProject := []*entity.Project{}
	err := r.db.Model(&entity.Project{}).
		Distinct().
		Select("project.id", "project.name", "project.description", "project.image", "project.created_at", "project.updated_at").
		Joins("join user_project_role upr on upr.project_id = project.id").
		Where("project.user_id = ?", userId).
		Where("project.status_id = ?", statusId).
		Offset(offset).
		Limit(size).
		Order(fmt.Sprintf("%v %v", sortBy, sortType)).
		Find(&listProject).Error
	if err != nil {
		return nil, err
	}

	return listProject, nil
}

func (r ProjectRepository) CreateProject(data *entity.Project) error {
	err := r.db.Model(&entity.Project{}).Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (r ProjectRepository) GetListMemberByProject(projectId int, page, size int, keyword, sortType, sortBy string) ([]*entity.UserTaskCount, error){
	offset := util.CalculateOffset(page, size)
	if sortType == "" && sortBy == "" {
		sortType = "DESC"
		sortBy = "created_at"
	}
	if sortType == "" {
		sortType = "DESC"
	}

	switch sortBy {
	case "createdAt":
		sortBy = "created_at"

	case "updatedAt":
		sortBy = "updated_at"

	default:
		sortBy = "created_at"
	}

	listUser := []*entity.UserTaskCount{}
	chain := r.db.Model(&entity.User{}).
		Select(`"user".*, COUNT(DISTINCT t.id) AS task_count`).
		Preload("Role").
		Preload("Status").
		Joins(`left join user_project_role upr on upr.user_id = "user".id`).
		Joins("left join project p on p.id = upr.project_id").
		Joins(`left join task t on t.assignee_id = "user".id OR t.reviewer_id = "user".id`).
		Where("upr.project_id = ?", projectId)
		
	if keyword != "" {
		chain = chain.Where(`"user".username LIKE ?`, "%"+keyword+"%").
			Or(`"user".name LIKE ?`, "%"+keyword+"%")
	}
	
	err := chain.Group(`"user".id`).
		Offset(offset).
		Limit(size).
		Order(fmt.Sprintf("%v %v", sortBy, sortType)).
		Scan(&listUser).Error

	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (r ProjectRepository) CountListMemberByProject(projectId int) (int, error){
	var count int64
	err := r.db.Model(&entity.User{}).
		Select(`"user".*`).
		Preload("Role").
		Preload("Status").
		Joins(`left join user_project_role upr on upr.user_id = "user".id`).
		Joins("left join project p on p.id = upr.project_id").
		Joins(`left join task t on t.assignee_id = "user".id OR t.reviewer_id = "user".id`).
		Where("upr.project_id = ?", projectId).
		Group(`"user".id`).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r ProjectRepository) FindById(id int ) (*entity.Project, error) {
	project := &entity.Project{}
	err := r.db.Model(&entity.Project{}).
		Preload("Status").
		Where("id = ?", id).First(&project).Error

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (r ProjectRepository) CountListTaskOpenUser(projectId, userId, statusId int) (*entity.UserTaskCount, error) {
	userTaskCount := &entity.UserTaskCount{}
	err := r.db.Model(&entity.User{}).
		Select(`"user".*, COUNT(t.id) as task_count`).
		Joins(`join task t on t.assignee_id = "user".id`).
		Where("t.project_id = ?", projectId).
		Where("t.assignee_id = ?", userId).
		Where("t.status_id != ?", statusId).
		Group(`"user".id`).
		Scan(&userTaskCount).Error

	if err != nil {
		return nil, err
	}

	return userTaskCount, nil
}

func (r ProjectRepository) CountListTaskByStatus(projectId, userId, statusId int) (*entity.UserTaskCount, error) {
	userTaskCount := &entity.UserTaskCount{}
	err := r.db.Model(&entity.User{}).
		Select(`"user".*, COUNT(t.id) as task_count`).
		Joins(`join task t on t.assignee_id = "user".id`).
		Where("t.project_id = ?", projectId).
		Where("t.assignee_id = ?", userId).
		Where("t.status_id = ?", statusId).
		Group(`"user".id`).
		Scan(&userTaskCount).Error

	if err != nil {
		return nil, err
	}

	return userTaskCount, nil
}

func (r ProjectRepository) UpdateProject(projectId int, mapData map[string]interface{}) error {
	err := r.db.Model(&entity.Project{}).
		Where("id = ?", projectId).
		Updates(mapData).Error

	if err != nil {
		return err
	}

	return nil
}

func (r ProjectRepository) GetAllProject(userId, page, size int, sortType, sortBy string) ([]*entity.Project, error) {
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
		sortBy = "project.created_at"

	case "updatedAt":
		sortBy = "project.updated_at"

	default:
		sortBy = "project.created_at"
	}
	listProject := []*entity.Project{}
	err := r.db.Model(&entity.Project{}).
		Distinct().
		Preload("User").
		Preload("Status").
		Joins("join user_project_role upr on upr.project_id = project.id").
		Where("project.user_id = ?", userId).
		Offset(offset).
		Limit(size).
		Order(fmt.Sprintf("%v %v", sortBy, sortType)).
		Find(&listProject).Error
	if err != nil {
		return nil, err
	}

	return listProject, nil
}

func (r ProjectRepository) CountAllProject(userId int) (int, error) {
	var count int64
	err := r.db.Model(&entity.Project{}).
		Where("user_id = ?", userId).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r ProjectRepository) DeleteProject(listId []int) error {
	err := r.db.Delete(&entity.Project{}, listId).Error
	if err != nil {
		return err
	}

	return nil
}