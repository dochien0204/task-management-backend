package payload

type TaskPayload struct {
	Name string `json:"name"`
	Description string `json:"description"`
	AssigneeId int `json:"assigneeId"`
	ReviewerId int `json:"reviewerId"`
	CategoryId int `json:"categoryId"`
	ProjectId int `json:"projectId"`
	StartDate string `json:"startDate"`
	DueDate string `json:"dueDate"`
}

type TaskUpdatePayload struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	AssigneeId int `json:"assigneeId"`
	ReviewerId int `json:"reviewerId"`
	CategoryId int `json:"categoryId"`
	StatusId int `json:"statusId"`
	StartDate string `json:"startDate"`
	DueDate string `json:"dueDate"`
}