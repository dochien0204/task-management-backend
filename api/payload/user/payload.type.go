package payload

type UpdateAvatar struct {
	Avatar string `json:"avatar"`
}

type DeleteUser struct {
	ListUserId []int `json:"listUserId"`
}

type UserUpdatePayload struct {
	Id int `json:"id"`
	Name string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address string `json:"address"`
	Email string `json:"email"`
}

type UserChangePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}