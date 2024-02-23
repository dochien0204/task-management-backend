package handler

import (
	rolePresenter "source-base-go/api/presenter/role"
	"source-base-go/config"
	"source-base-go/entity"
)

func convertListRoleEntityToPresenter(listData []*entity.Role) []*rolePresenter.Role {
	listRolePresenter := []*rolePresenter.Role{}
	for _, role := range listData {
		rolePresenter := &rolePresenter.Role{
			Name:        role.Name,
			Code:        role.Code,
			Description: role.Description,
			CreatedAt:   role.CreatedAt.Format(config.LAYOUT),
			UpdatedAt:   role.UpdatedAt.Format(config.LAYOUT),
		}
		listRolePresenter = append(listRolePresenter, rolePresenter)
	}

	return listRolePresenter
}

func convertRoleEntityToPresenter(data *entity.Role) *rolePresenter.Role {
	return &rolePresenter.Role{
		Name:        data.Name,
		Code:        data.Code,
		Description: data.Description,
		CreatedAt:   data.CreatedAt.Format(config.LAYOUT),
		UpdatedAt:   data.UpdatedAt.Format(config.LAYOUT),
	}
}
