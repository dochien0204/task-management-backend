package presenter

import statusPresenter "source-base-go/api/presenter/status"

type UserProfilePresenter struct {
	Id       int                              `json:"id"`
	Username string                           `json:"userName"`
	Status   *statusPresenter.StatusPresenter `json:"status"`
}

type UserProfilePresenterResponse struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Result  *UserProfilePresenter `json:"result"`
}
