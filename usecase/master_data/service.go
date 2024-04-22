package masterdata

import "source-base-go/entity"

type Service struct {
	statusRepo StatusRepository
}

func NewService(statusRepo StatusRepository) *Service {
	return &Service{
		statusRepo: statusRepo,
	}
}

func (s Service) FindStatusByType(typeStatus string) ([]*entity.Status, error) {
	return s.statusRepo.FindByType(typeStatus)
}