package project

import (
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"

	"gorm.io/gorm"
)

type Service struct {
	projectRepo         ProjectRepository
	userProjectRoleRepo UserProjectRoleRepository
	roleRepo            RoleRepository
}

func NewService(projectRepo ProjectRepository, userProjectRoleRepo UserProjectRoleRepository, roleRepo RoleRepository) *Service {
	return &Service{
		projectRepo:         projectRepo,
		userProjectRoleRepo: userProjectRoleRepo,
		roleRepo:            roleRepo,
	}
}

func (s Service) WithTrx(trxHandle *gorm.DB) Service {
	s.projectRepo = s.projectRepo.WithTrx(trxHandle)
	s.userProjectRoleRepo = s.userProjectRoleRepo.WithTrx(trxHandle)

	return s
}

func (s Service) CreateProject(userId int, project *entity.Project) error {
	//Check project name is exists
	projectCreate, err := s.projectRepo.FindProjectByName(project.Name)
	if err != nil {
		return err
	}

	if projectCreate.Name != "" {
		return entity.ErrProjectAlreadyExists
	}

	//Create project
	err = s.projectRepo.CreateProject(project)
	if err != nil {
		return err
	}

	//Create user project role owner for user
	projectId := project.Id

	roleOwnerProject, err := s.roleRepo.FindByCode(string(define.OWNER), string(define.PROJECT))
	if err != nil {
		return err
	}

	userProjectRole := &entity.UserProjectRole{
		UserId:    userId,
		RoleId:    roleOwnerProject.Id,
		ProjectId: projectId,
	}

	err = s.userProjectRoleRepo.CreateUserProjectRole(userProjectRole)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetListProjectOfUser(userId, page, size int, sortType, sortBy string) ([]*entity.Project, error) {
	return s.projectRepo.GetListProjectOfUser(userId, page, size, sortType, sortBy)
}

func (s Service) AddListMemberToProject(userId, projectId int, listUserId []int) error {
	listUserProjectRole := []*entity.UserProjectRole{}

	//Check user is owner of project
	roleOwner, err := s.roleRepo.FindByCode(string(define.OWNER), string(define.PROJECT))
	if err != nil {
		return err
	}

	userProjectOwner, err := s.userProjectRoleRepo.GetProjectOwner(projectId, roleOwner.Id)
	if err != nil {
		return err
	}

	if userProjectOwner.UserId != userId {
		return entity.ErrForbidden
	}

	//Get role member project
	roleMemberProject, err := s.roleRepo.FindByCode(string(define.MEMBER), string(define.PROJECT))
	if err != nil {
		return err
	}

	for _, userId := range listUserId {
		userProjectRole := &entity.UserProjectRole{
			ProjectId: projectId,
			UserId:    userId,
			RoleId:    roleMemberProject.Id,
		}

		listUserProjectRole = append(listUserProjectRole, userProjectRole)
	}

	//Create list user projectt role
	err = s.userProjectRoleRepo.CreateListUserProjectRole(listUserProjectRole)
	if err != nil {
		return err
	}

	return nil
}
