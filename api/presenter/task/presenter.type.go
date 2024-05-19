package presenter

type ListTaskPresenter struct {
	Status *Status `json:"status"`
	ListTask []*Task `json:"listTask"`
}

type Task struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Category *Category `json:"category"`
	Assignee *UserPresenter `json:"assignee"`
}

type Category struct {
	Id int `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Status struct {
	Id          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type TaskDetail struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"string"`
	Category *Category `json:"category"`
	User *User `json:"createdBy"`
	Assignee *User `json:"assignee"`
	Reviewer *User `json:"reviewer"`
	Documents []*Document `json:"documents"`
	StartDate string `json:"startDate"`
	Status *Status `json:"status"`
	DueDate string `json:"dueDate"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Document struct {
	Id int `json:"id"`
	Name string `json:"name"`
	FileName string `json:"fileName"`
	TaskId int `json:"taskId"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type User struct {
	Id       int `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type ListTaskByDatePresenter struct {
	Date string `json:"date"`
	ListTask []*TaskDetail `json:"listTask"`
	Count int `json:"count"`
}

type DiscussionPresenter struct {
	Id int `json:"id"`
	Comment string `json:"comment"`
	TaskId int
	User *UserPresenter `json:"user"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserPresenter struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
}

type ListTaskByUserAndStatusPresenter struct {
	ListTask []*TaskDetail `json:"listTask"`
}