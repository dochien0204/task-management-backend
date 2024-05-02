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