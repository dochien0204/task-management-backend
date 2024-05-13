package payload

type ProjectPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ListUserIdProjectPayload struct {
	ProjectId  int   `json:"projectId"`
	ListUserId []int `json:"listUserId"`
	RoleId int `json:"roleId"`
}

type UserOverViewPayload struct {
	ProjectId int `json:"projectId"`
	UserId int `json:"userId"`
}

type ProjectUpdatePayload struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Name string `json:"name"`
	StatusId int `json:"statusId"`
}

type ProjectDeletePayload struct {
	ListProjectId []int `json:"listProjectId"`
}