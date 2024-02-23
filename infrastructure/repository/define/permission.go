package define

type Permisstion string
type RoleType string

const (
	ADMIN  Permisstion = "admin"
	USER   Permisstion = "user"
	SYSTEM RoleType    = "system"
)
