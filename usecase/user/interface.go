package user

import (
	payload "source-base-go/api/payload/user"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository"

	"gorm.io/gorm"
)

type Verifier interface {
	CacheUserData(user *entity.User, listRole []string, expiredAt int) error
}

type UserRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.UserRepository

	GetUserProfile(userId int) (*entity.User, error)
	FindByUsername(userName string) (*entity.User, error)
	CreateUser(user *entity.User) error
	IsUserExists(username string) (bool, error)
	GetListUser(statusId, page, size int, sortType, sortBy string) ([]*entity.User, error)
	CountListUser(statusId int) (int, error)
	IsUserEmailExists(email string) (bool, error)
	UpdateAvatar(userId int, avatar string) error
	DeleteUserById(id []int) error
	UpdateUser(userId int, mapData map[string]interface{}) error
	FindById(id int ) (*entity.User, error)
	ChangePassword(userId int , password string) error
}

type RoleRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.RoleRepository

	FindAllRolesOfUser(userId int) ([]*entity.Role, error)
	FindByCode(code string, typeRole string) (*entity.Role, error)
}

type UserRoleRepository interface {
	WithTrx(trxHandle *gorm.DB) repository.UserRoleRepository

	AddRoleForUser(userId, roleId int) error
}

type StatusRepository interface {
	GetStatusByCodeAndType(typeStatus string, code string) (*entity.Status, error)
}

type EmailRepository interface {
	SendMailPasswordForUser(body string, to []string, subject string) error
}

type UseCase interface {
	WithTrx(trxHandle *gorm.DB) Service

	GetUserProfile(userId int) (*entity.User, error)
	Login(username string, password string) (*entity.TokenPair, *entity.User, error)
	Register(user *entity.User) error
	GetListUser(page, size int, sortType, sortBy string) ([]*entity.User, int, error)
	UpdateAvatar(userId int, avatar string) error
	DeleteUserById(id []int) error
	ChangePassword(userId int, payload payload.UserChangePassword) error
	ResetPassword(userId int) error
}
