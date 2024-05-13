package project

import (
	payload "source-base-go/api/payload/project"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	projectRepo         ProjectRepository
	userProjectRoleRepo UserProjectRoleRepository
	roleRepo            RoleRepository
	activityRepo ActivityRepository
	statusRepo StatusRepository
}

func NewService(projectRepo ProjectRepository, userProjectRoleRepo UserProjectRoleRepository, roleRepo RoleRepository, activityRepo ActivityRepository, statusRepo StatusRepository) *Service {
	return &Service{
		projectRepo:         projectRepo,
		userProjectRoleRepo: userProjectRoleRepo,
		roleRepo:            roleRepo,
		activityRepo: activityRepo,
		statusRepo: statusRepo,
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

	statusActive, err := s.statusRepo.GetStatusByCodeAndType(define.PROJECT_CODE, define.PROJECT_ACTIVE_CODE)
	if err != nil {
		return err
	}

	project.StatusId = statusActive.Id
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
	statusProjectActive, err := s.statusRepo.GetStatusByCodeAndType(define.PROJECT_CODE, define.PROJECT_ACTIVE_CODE)
	if err != nil {
		return nil, err
	}
	return s.projectRepo.GetListProjectOfUser(userId, statusProjectActive.Id, page, size, sortType, sortBy)
}

func (s Service) AddListMemberToProject(userId, projectId, roleId int, listUserId []int) error {
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

	for _, userId := range listUserId {
		userProjectRole := &entity.UserProjectRole{
			ProjectId: projectId,
			UserId:    userId,
			RoleId:    roleId,
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

func (s Service) GetProjectDetail(projectId int) (*entity.UserProjectRole, error) {
	//Get role owner
	roleOwnerProject, err := s.roleRepo.FindByCode(string(define.OWNER), string(define.PROJECT))
	if err != nil {
		return nil, err
	}

	projectDetail, err := s.userProjectRoleRepo.GetProjectDetailWithOwner(projectId, roleOwnerProject.Id)
	if err != nil {
		return nil, err
	}

	return projectDetail, nil
}

func (s Service) GetListMemberByProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.UserTaskCount, int, error) {
	listMember, err := s.projectRepo.GetListMemberByProject(projectId, page, size, sortType, sortBy)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.projectRepo.CountListMemberByProject(projectId)
	if err != nil {
		return nil, 0, err
	}

	return listMember, count, nil
}

func (s Service) GetListActivityProjectByDate(projectId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Activity, error) {
	return s.activityRepo.GetListActivityByDate(projectId, timeOffset, fromDate, toDate)
}

func (s Service) GetOverviewUserTaskProject(projectId, userId int) (*entity.UserTaskCount, *entity.UserTaskCount, *entity.UserProjectRole, error) {
	//Get status closed task
	statusClosedTask, err := s.statusRepo.GetStatusByCodeAndType(define.TASK_CODE, define.TASK_CLOSED_STATUS)
	if err != nil {
		return nil, nil, nil, err
	}

	//Get role user in project
	userProjectRole, err := s.userProjectRoleRepo.GetRoleUserInProject(projectId, userId)
	if err != nil {
		return nil, nil, nil, err
	}

	//Get task not in closed status
	listTaskNotClosed, err := s.projectRepo.CountListTaskOpenUser(projectId, userId, statusClosedTask.Id)
	if err != nil {
		return nil, nil, nil, err
	}

	listTaskClosed, err := s.projectRepo.CountListTaskByStatus(projectId, userId, statusClosedTask.Id)
	if err != nil {
		return nil, nil, nil, err
	}

	return listTaskNotClosed, listTaskClosed, userProjectRole, nil
}

func (s Service) GetListActivityByDateOfUser(projectId, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Activity, error) {
	return s.activityRepo.GetListActivityByDateOfUser(projectId, userId, timeOffset, fromDate, toDate)
} 

func (s Service) UpdateProject(payload payload.ProjectUpdatePayload) error {
	projectUpdate, err := s.projectRepo.FindById(payload.Id)
	if err != nil {
		return err
	}

	if payload.Name != projectUpdate.Name {
		projectCreate, err := s.projectRepo.FindProjectByName(payload.Name)
		if err != nil {
			return err
		}

		if projectCreate.Name != "" {
			return entity.ErrProjectAlreadyExists
		}
	}

	mapData := map[string]interface{}{}
	mapData["name"] = payload.Name
	mapData["description"] = payload.Description
	mapData["status_id"] = payload.StatusId

	err = s.projectRepo.UpdateProject(payload.Id, mapData)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetAllProject(userId, page, size int, sortType, sortBy string) ([]*entity.Project, int, error) {
	count, err := s.projectRepo.CountAllProject(userId)
	if err != nil {
		return nil, 0, err
	}

	listProject, err := s.projectRepo.GetAllProject(userId, page, size, sortType, sortBy)
	if err != nil {
		return nil, 0, err
	}

	return listProject, count, nil
}

func (s Service) DeleteProjectByListId(listId []int) error {
	return s.projectRepo.DeleteProject(listId)
}