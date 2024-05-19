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
	Role *Role `json:"category"`
}

type ListActivityProjectByDate struct {
	Date string `json:"date"`
	ListActivity []*Activity `json:"listActivity"`
}

type Activity struct {
	Id int `json:"id"`
	TaskId int `json:"taskId"`
	User *UserPresenter `json:"user"`
	Description string `json:"description"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserProjectOverview struct {
	User *UserPresenter `json:"user"`
	TaskOpenCount int `json:"taskOpenCount"`
	TaskCloseCount int `json:"taskClosedCount"`
}

type Role struct {
	Id int `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ProjectAdminPresenter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	User *UserPresenter `json:"user"`
	Status *Status `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type Status struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type ProjectChartOverview struct {
	Project *ProjectPresenter `json:"project"`
	TotalTask int `json:"totalTask"`
	DoneTask int `json:"doneTask"`
	MemberCount int `json:"memberCount"`
}