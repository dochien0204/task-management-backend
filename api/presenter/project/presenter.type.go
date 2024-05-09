package presenter

type ProjectPresenter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type ProjectDetailPresenter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Owner       *User  `json:"owner"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type UserTaskCount struct {
	User *UserPresenter `json:"user"`
	TaskCount int `json:"taskCount"`
}

type UserPresenter struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
}

