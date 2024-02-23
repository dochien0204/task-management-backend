package payload

type ProjectPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ListUserIdProjectPayload struct {
	ProjectId  int   `json:"projectId"`
	ListUserId []int `json:"listUserId"`
}
