package masterdata

import "source-base-go/entity"

type Service struct {
	statusRepo StatusRepository
	categoryRepo CategoryRepository
}

func NewService(statusRepo StatusRepository, categoryRepo CategoryRepository) *Service {
	return &Service{
		statusRepo: statusRepo,
		categoryRepo: categoryRepo,
	}
}

func (s Service) FindStatusByType(typeStatus string) ([]*entity.Status, error) {
	return s.statusRepo.FindByType(typeStatus)
}

func (s Service) GetListCategoryByCode(typeCate string) ([]*entity.Category, error) {
	return s.categoryRepo.GetListCategoryByCode(typeCate)
}