package payload

type TaskPayload struct {
	Name string `json:"name"`
	Description string `json:"description"`
	StatusId int `json:"statusId"`
	ProjectId int `json:"projectId"`
}

type TaskDocumentPayload struct {
	File string `json:"file"`
	Name string `json:"name"`
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
	Documents []*TaskDocumentPayload `json:"documents"`
}

type TaskStatusUpdatePayload struct {
	Id int `json:"id"`
	StatusId int `json:"statusId"`
}