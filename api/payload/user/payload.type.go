package payload

type UpdateAvatar struct {
	Avatar string `json:"avatar"`
}

type DeleteUser struct {
	ListUserId []int `json:"listUserId"`
}