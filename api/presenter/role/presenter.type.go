package presenter

type Role struct {
	Id int `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type ListRoleResp struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Results []*Role `json:"results"`
}

type RoleResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Results *Role  `json:"results"`
}
