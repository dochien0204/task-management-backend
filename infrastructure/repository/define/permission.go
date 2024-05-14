package define

type Permisstion string
type RoleType string

const (
	ADMIN  Permisstion = "admin"
	USER   Permisstion = "user"
	OWNER  Permisstion = "owner"
	MEMBER Permisstion = "member"
	PROJECT_MANAGER Permisstion = "manager"

	PROJECT RoleType = "project"
	SYSTEM  RoleType = "system"
)
