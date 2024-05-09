package repository

import (
	"source-base-go/entity"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r CategoryRepository) GetListCategoryByCode(typeCate string) ([]*entity.Category, error) {
	listCate := []*entity.Category{}
	err := r.db.Model(&entity.Category{}).
		Where("type = ?", typeCate).
		Find(&listCate).Error

	if err != nil {
		return nil, err
	}

	return listCate, nil
}