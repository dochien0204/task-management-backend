package presenter

import statusPresenter "source-base-go/api/presenter/status"

type UserProfilePresenter struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Address string `json:"address"`
	Status   *statusPresenter.StatusPresenter `json:"status"`
}

type UserProfilePresenterResponse struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Result  *UserProfilePresenter `json:"result"`
}

type UserPresenter struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
}
