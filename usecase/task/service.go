package task

import (
	"fmt"
	taskPayload "source-base-go/api/payload/task"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	taskRepo TaskRepository
	statusRepo StatusRepository
	taskDocumentRepo TaskDocumentRepository
	discussionRepo DiscussionRepository
	activityRepo ActivityRepository
	userRepo UserRepository
	projectRepo ProjectRepository
	userProjectRoleRepo UserProjectRoleRepository
	emailRepo EmailRepository
}

func NewService(taskRepo TaskRepository, statusRepo StatusRepository, taskDocumentRepo TaskDocumentRepository, discussionRepo DiscussionRepository, activityRepo ActivityRepository, userRepo UserRepository, projectRepo ProjectRepository, userProjectRoleRepo UserProjectRoleRepository,
	emailRepo EmailRepository) *Service {
	return &Service{
		taskRepo: taskRepo,
		statusRepo: statusRepo,
		taskDocumentRepo: taskDocumentRepo,
		discussionRepo: discussionRepo,
		activityRepo: activityRepo,
		userRepo: userRepo,
		projectRepo: projectRepo,
		userProjectRoleRepo: userProjectRoleRepo,
		emailRepo: emailRepo,
	}
}

func (s Service) WithTrx(trxHandle *gorm.DB) Service {
	s.taskRepo = s.taskRepo.WithTrx(trxHandle)
	s.taskDocumentRepo = s.taskDocumentRepo.WithTrx(trxHandle)
	s.discussionRepo = s.discussionRepo.WithTrx(trxHandle)
	s.activityRepo = s.activityRepo.WithTrx(trxHandle)

	return s
}

func (s Service) CreateTask(userId int, payload taskPayload.TaskPayload) (int, error) {
	data := &entity.Task {
		Name: payload.Name,
		Description: payload.Description,
		ProjectId: payload.ProjectId,
		UserId: userId,
		StatusId: payload.StatusId,
		AssigneeId: payload.AssigneeId,
		ReviewerId: payload.ReviewerId,
		CategoryId: payload.CategoryId,
	}

	if payload.StartDate != nil {
		startDate, _ := time.Parse(config.LAYOUT, *payload.StartDate)
		data.StartDate = &startDate
	}

	if payload.DueDate != nil {
		dueDate, _ := time.Parse(config.LAYOUT, *payload.DueDate)
		data.DueDate = &dueDate
	}

	err := s.taskRepo.Create(data)
	if err != nil {
		return 0, err
	}

	//Create task document
	if len(payload.Documents) > 0 {
		listTaskDocument := []*entity.TaskDocument{}
		for _, documentPayload := range payload.Documents {
			document := &entity.TaskDocument{
				File: documentPayload.File,
				Name: documentPayload.Name,
				TaskId: data.Id,
			}
	
			listTaskDocument = append(listTaskDocument, document)
		}
	
		err = s.taskDocumentRepo.CreateDocumentsForTask(listTaskDocument)
		if err != nil {
			return 0, err
		}
	}

	//Create activity
	user, err := s.userRepo.FindById(userId)
	if err != nil {
		return 0, err
	}

	project, err := s.projectRepo.FindById(payload.ProjectId)
	if err != nil {
		return 0, err
	}

	description := fmt.Sprintf("User (%v) has created a task (%v) for the project (%v)", user.Username, data.Name, project.Name)
	activity := &entity.Activity{
		UserId: userId,
		TaskId: data.Id,
		Description: description,
	}

	// //Find all user in project
	// listUser, err := s.userProjectRoleRepo.FindAllUserOfProject(data.ProjectId)
	// if err != nil {
	// 	return 0, err
	// }

	// listEmail := []string{}
	// for _, user := range listUser {
	// 	listEmail = append(listEmail, user.Email)
	// }

	// err = s.emailRepo.SendMailForUsers(description, listEmail, "TASK CREATE")
	// if err != nil {
	// 	return 0, err
	// }

	//Send mail
	err = s.activityRepo.CreateActivity(activity)
	if err != nil {
		return 0, err
	}
	
	return data.Id, nil
}

func (s Service) GetListTaskOfProject(projectId int, page, size int, sortType, sortBy string) ([]*entity.Task, []*entity.Status, error) {
	listStatusTask, err := s.statusRepo.FindByType(define.TASK_CODE)
	if err != nil {
		return nil, nil, err
	}

	listTask, err := s.taskRepo.GetListTaskOfProject(projectId, page, size, sortBy, sortType)
	if err != nil {
		return nil, nil, err
	}

	return listTask, listStatusTask, nil 
}

func (s Service) GetTaskDetail(taskId int) (*entity.Task, error) {
	return s.taskRepo.GetTaskDetail(taskId)
}

func (s Service) UpdateTask(userId int, data taskPayload.TaskUpdatePayload) error {
	mapTask := map[string]interface{}{}
	mapTask["name"] = data.Name
	mapTask["description"] = data.Description
	mapTask["assignee_id"] = data.AssigneeId
	mapTask["reviewer_id"] = data.ReviewerId
	mapTask["category_id"] = data.CategoryId
	mapTask["status_id"] = data.StatusId

	if data.StartDate != nil {
		startDate, _ := time.Parse(config.LAYOUT, *data.StartDate)
		mapTask["start_date"]= startDate
	} else {
		mapTask["start_date"] = data.StartDate
	}
	
	if data.DueDate != nil {
		dueDate, _ := time.Parse(config.LAYOUT, *data.DueDate)
		mapTask["due_date"] = dueDate
	} else {
		mapTask["due_date"] = data.DueDate
	}
	task, err := s.taskRepo.GetTaskDetail(data.Id)
	if err != nil {
		return err
	}
	
	userLogg, err := s.userRepo.FindById(userId)
	if err != nil {
		return err
	}

	assignee, err := s.userRepo.FindById(*data.AssigneeId)
	if err != nil {
		return err
	}

	description := ""
	if task.Name != data.Name {
		description += fmt.Sprintf("Người dùng (%v) đã thay đổi tên task (%v) thành (%v) |", userLogg.Username, task.Name, data.Name)
	}

	if task.Description != data.Description {
		description += fmt.Sprintf("Người dùng (%v) đã thay đổi mô tả cho task (%v) thành (%v) |", userLogg.Username, task.Name, data.Description)
	}

	if *task.AssigneeId != *data.AssigneeId {
		description += fmt.Sprintf("Người dùng (%v) đã thay đổi người nhận task (%v) từ (%v) sang người dùng (%v) | ", userLogg.Username, task.Name,task.Assignee.Username, assignee.Username)
	}
	err = s.taskRepo.UpdateTask(data.Id, mapTask)
	if err != nil {
		return err
	}

	activity := &entity.Activity{
		UserId: userId,
		TaskId: data.Id,
		Description: description,
	}

	// //Find all user in project
	// listUser, err := s.userProjectRoleRepo.FindAllUserOfProject(data.ProjectId)
	// if err != nil {
	// 	return 0, err
	// }

	// listEmail := []string{}
	// for _, user := range listUser {
	// 	listEmail = append(listEmail, user.Email)
	// }

	// err = s.emailRepo.SendMailForUsers(description, listEmail, "TASK CREATE")
	// if err != nil {
	// 	return 0, err
	// }

	//Send mail
	if activity.Description != "" {
		err = s.activityRepo.CreateActivity(activity)
		if err != nil {
			return err
		}
	}

	//Create task document
	listTaskDocument := []*entity.TaskDocument{}
	//Delete all task of document
	err = s.taskDocumentRepo.DeleteAllTaskDocumentOfTask(data.Id)
	if err != nil {
		return err
	}

	for _, documentPayload := range data.Documents {
		document := &entity.TaskDocument{
			File: documentPayload.File,
			Name: documentPayload.Name,
			TaskId: data.Id,
		}

		listTaskDocument = append(listTaskDocument, document)
	}

	if len(data.Documents) > 0 {
		err = s.taskDocumentRepo.CreateDocumentsForTask(listTaskDocument)
		if err != nil {
			return err
		}
	}
	
	return nil
}

func (s Service) UpdateTaskStatus(userId, taskId, statusId int) error {
	status, err := s.statusRepo.FindById(statusId)
	if err != nil {
		return err
	}

	task, err := s.taskRepo.GetTaskDetail(taskId)
	if err != nil {
		return err
	}
	
	userLogg, err := s.userRepo.FindById(userId)
	if err != nil {
		return err
	}

	description := ""
	if statusId != task.StatusId {
		description += fmt.Sprintf("Người dùng (%v) đã thay đổi người trạng task (%v) từ (%v) sang trạng thái (%v) | ", userLogg.Username, task.Name, task.Status.Name, status.Name)
	}

	err = s.taskRepo.UpdateStatusTask(taskId, statusId)
	if err != nil {
		return err
	}

	activity := &entity.Activity{
		UserId: userId,
		TaskId: taskId,
		Description: description,
	}

	//Send mail
	if activity.Description != "" {
		err = s.activityRepo.CreateActivity(activity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) GetListTaskByDate(projectId int, userId int, timeOffset int, fromDate time.Time, toDate time.Time) ([]*entity.Task, error) {
	return s.taskRepo.GetListTaskByDate(projectId, userId, timeOffset, fromDate, toDate)
}

func (s Service) CreateDiscussionTask(userId, taskId int, comment string) error {
	discussion := &entity.Discussion{
		UserId: userId,
		TaskId: taskId,
		Comment: comment,
	}

	return s.discussionRepo.Create(discussion)
}

func (s Service) GetListDiscussionOfTask(taskId int, page, size int, sortBy, sortType string) ([]*entity.Discussion, int, error) {
	count, err := s.discussionRepo.CountListDiscussion(taskId)
	if err != nil {
		return nil, 0, err
	}

	listDiscussion, err := s.discussionRepo.GetListDiscussion(taskId, page,size, sortBy, sortType)
	if err != nil {
		return nil, 0, err
	}
	return listDiscussion, count, nil
}

func (s Service) GetListTaskProjectByUserAndStatus(projectId int, assigneeId, statusId int, page, size int, sortType, sortBy string) ([]*entity.Task, int, error) {


	if statusId != 0 {
		listTask, err := s.taskRepo.GetListTaskProjectByUserAndStatus(projectId, assigneeId, statusId, page, size, sortType, sortBy)
		if err != nil {
			return nil, 0, err
		}
	
		count, err := s.taskRepo.CountListTaskProjectByUserAndStatus(projectId, assigneeId, statusId)
		if err != nil {
			return nil, 0, err
		}

		return listTask, count, nil
	}

	listTask, err := s.taskRepo.GetListTaskProjectByUser(projectId, assigneeId, page, size, sortType, sortBy)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.taskRepo.CountListTaskProjectByUser(projectId, assigneeId)
	if err != nil {
		return nil, 0, err
	}

	return listTask, count, nil
}

func (s Service) DeleteTask(userId, taskId int) error {
	task, err := s.taskRepo.GetTaskDetail(taskId)
	if err != nil {
		return err
	}

	roles, err := s.userProjectRoleRepo.GetRoleUserInProject(task.ProjectId, userId)
	if err != nil {
		return err
	}

	if roles.Role.Code != string(define.OWNER) && roles.Role.Code != string(define.PROJECT_MANAGER) {
		return entity.ErrNotHavePermissionDeleteTask
	}

	err = s.taskRepo.DeleteTask(taskId)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetListTaskByUser(userId int, keyword string, page, size int, sortBy, sortType string) ([]*entity.Task, int, error) {
	//Get list project user join
	listUserProjectRole, err := s.userProjectRoleRepo.GetAllProjectOfUser(userId)
	if err != nil {
		return nil, 0, err
	}

	listProjectId := []int{}
	for _, userProjectRole := range listUserProjectRole {
		listProjectId = append(listProjectId, userProjectRole.ProjectId)
	}

	//Find list task by list project id
	listTask, err := s.taskRepo.GetListTaskByUser(listProjectId, keyword, page, size, sortBy, sortType)
	if err != nil {
		return nil, 0, err
	}

	//Count
	count, err := s.taskRepo.CountListTaskByUser(listProjectId, keyword)
	if err != nil {
		return nil, 0, err
	}

	return listTask, count, nil
}