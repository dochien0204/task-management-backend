package entity

type UserProjectRole struct {
	ProjectId int
	UserId    int
	RoleId    int
	Project   *Project
	User      *User
	Role *Role
}
