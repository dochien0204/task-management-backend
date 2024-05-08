package user

import (
	"fmt"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"
	"source-base-go/infrastructure/repository/util"

	"gorm.io/gorm"
)

type Service struct {
	userRepository UserRepository
	roleRepo       RoleRepository
	userRoleRepo   UserRoleRepository
	StatusRepo 	StatusRepository
	verifier       Verifier
}

func NewService(userRepository UserRepository, roleRepo RoleRepository, userRoleRepo UserRoleRepository, statusRepo StatusRepository, verifier Verifier) *Service {
	return &Service{
		userRepository: userRepository,
		roleRepo:       roleRepo,
		userRoleRepo:   userRoleRepo,
		StatusRepo: statusRepo,
		verifier:       verifier,
	}
}

func (s Service) WithTrx(trxHandle *gorm.DB) Service {
	s.userRepository = s.userRepository.WithTrx(trxHandle)
	s.roleRepo = s.roleRepo.WithTrx(trxHandle)
	s.userRoleRepo = s.userRoleRepo.WithTrx(trxHandle)
	return s
}

func (s *Service) GetUserProfile(userId int) (*entity.User, error) {
	return s.userRepository.GetUserProfile(userId)
}

func (s *Service) Login(username string, password string) (*entity.TokenPair, *entity.User, error) {
	//Get user by user name
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}

	if user == nil {
		return nil, nil, entity.ErrUsernameNotExists
	}

	//Vefiry password
	passwordValidate := user.ValidatePassword(password)
	if !passwordValidate {
		return nil, nil, entity.ErrInvalidPassword
	}

	//Find all roles of user
	listRole, err := s.roleRepo.FindAllRolesOfUser(user.Id)
	listRoleStr := []string{}
	for _, role := range listRole {
		listRoleStr = append(listRoleStr, role.Code)
	}
	//Cache userdata into redis
	err = s.verifier.CacheUserData(user, listRoleStr, config.GetInt("jwt.accessMaxAge"))
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}

	//Generate token
	acessToken, err := util.GenerateAccessToken(user)
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}
	refreshToken, err := util.GenerateRefreshToken(user)
	if err != nil {
		return nil, nil, entity.ErrInternalServerError
	}

	tokenPair := &entity.TokenPair{
		Token:        acessToken,
		RefreshToken: refreshToken,
	}

	return tokenPair, user, nil
}

func (s Service) Register(user *entity.User) error {
	//Check user exists
	isExists, err := s.userRepository.IsUserExists(user.Username)
	if err != nil {
		return entity.ErrBadRequest
	}

	if isExists {
		return entity.ErrAccountAlreadyExists
	}

	//Create user
	err = s.userRepository.CreateUser(user)
	if err != nil {
		return entity.ErrInternalServerError
	}

	//Find default role user
	roleDefault, err := s.roleRepo.FindByCode(string(define.USER), string(define.SYSTEM))
	if err != nil {
		return entity.ErrInternalServerError
	}
	//Add role for user
	err = s.userRoleRepo.AddRoleForUser(user.Id, roleDefault.Id)
	if err != nil {
		return entity.ErrInternalServerError
	}

	return nil
}

func (s Service) GetListUser(page, size int, sortType, sortBy string) ([]*entity.User, int, error) {
	statusUserActive, err := s.StatusRepo.GetStatusByCodeAndType(string(define.USER), string(define.ACTIVE))
	if err != nil {
		return nil, 0, err
	}

	listUser, err := s.userRepository.GetListUser(statusUserActive.Id, page, size, sortType, sortBy)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.userRepository.CountListUser(statusUserActive.Id)
	if err != nil {
		return nil, 0, err
	}

	return listUser, count, nil 
}

func (s Service) UpdateAvatar(userId int, avatar string) error {
	if avatar != "" {
		path := fmt.Sprintf("user/%d/%v", userId, avatar)
		avatarPath := fmt.Sprintf("https://%v.s3.%v.amazon.com/%v", config.S3_BUCKET_NAME, config.REGION, path)

		err := s.userRepository.UpdateAvatar(userId, avatarPath)
		if err != nil {
			return err
		}
	} else {
		err := s.userRepository.UpdateAvatar(userId, avatar)
		if err != nil {
			return err
		}
	}
	
	return nil
}
