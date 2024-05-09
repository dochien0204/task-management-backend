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

func (r ProjectRepository) GetListProjectOfUser(userId, page, size int, sortType, sortBy string) ([]*entity.Project, error) {
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
		Select("project.id", "project.name", "project.description", "project.image", "project.created_at", "project.updated_at").
		Joins("join user_project_role upr on upr.project_id = project.id").
		Where("user_id = ?", userId).
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

// func (r ProjectRepository) GetListMemberByProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.User, error){
// 	offset := util.CalculateOffset(page, size)
// 	if sortType == "" && sortBy == "" {
// 		sortType = "DESC"
// 		sortBy = "created_at"
// 	}
// 	if sortType == "" {
// 		sortType = "DESC"
// 	}

// 	switch sortBy {
// 	case "createdAt":
// 		sortBy = "created_at"

// 	case "updatedAt":
// 		sortBy = "updated_at"

// 	default:
// 		sortBy = "created_at"
// 	}

// 	listUser := []*entity.User{}
// 	err := r.db.Model(&entity.User{}).
// 		Select("u.*,")
// 		Joins("join user_project_role upr on upr.user_id = user.id").
// 		Where("upr.project_id = ?", projectId).
// 		Offset(offset).
// 		Limit(size).
// 		Order(fmt.Sprint("%v %v", sortBy, sortType)).
// 		Find(&listUser).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return listUser, nil
// }
