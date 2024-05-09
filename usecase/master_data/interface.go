package masterdata

import "source-base-go/entity"

type StatusRepository interface {
	FindByType(typeStatus string) ([]*entity.Status, error)
}

type CategoryRepository interface {
	GetListCategoryByCode(typeCate string) ([]*entity.Category, error)

}

type UseCase interface {
	FindStatusByType(typeStatus string) ([]*entity.Status, error)
	GetListCategoryByCode(typeCate string) ([]*entity.Category, error)
}