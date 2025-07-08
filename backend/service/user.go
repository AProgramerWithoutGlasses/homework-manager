package service

import (
	"backend/middleware"
	"backend/models"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login 老师登录
func (s *Service) Login(req *models.LoginRequest) (res *models.LoginResponse, err error) {
	// 获取老师
	teacher, err := s.dao.GetTeacherByTelephone(req.Telephone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("账户不存在")
		}
		return nil, fmt.Errorf("s.dao.GetTeacherByTelephone() err: %v", err)
	}

	// 验证密码
	if !s.CheckPassword(teacher.Password, req.Password) {
		return nil, errors.New("密码错误")
	}

	// 生成JWT令牌
	token, err := middleware.GenToken(teacher.ID, teacher.Telephone)
	if err != nil {
		return nil, fmt.Errorf("middleware.GenToken() err: %v", err)
	}

	res = &models.LoginResponse{
		Token:   token,
		Teacher: &teacher,
	}

	return
}

// CheckPassword 验证密码
func (s *Service) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// HashPassword 加密密码
func (s *Service) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// GetStudents 获取学生列表
func (s *Service) GetStudents() (students []models.Student, err error) {
	return s.dao.GetStudents()
}

// CreateStudent 创建学生
func (s *Service) CreateStudent(req *models.CreateStudentRequest) error {
	// 加密密码
	hashedPassword, err := s.HashPassword(req.Password)
	if err != nil {
		return err
	}

	student := &models.Student{
		Username: req.Username,
		Password: hashedPassword,
		Name:     req.Name,
		Class:    req.Class,
		Grade:    req.Grade,
	}

	return s.dao.CreateStudent(student)
}

// UpdateStudent 更新学生信息
func (s *Service) UpdateStudent(id uint, req *models.UpdateStudentRequest) error {
	student, err := s.dao.GetStudentByID(id)
	if err != nil {
		return err
	}

	student.Username = req.Username
	student.Name = req.Name
	student.Class = req.Class
	student.Grade = req.Grade

	// 如果提供了新密码，则更新密码
	if req.Password != "" {
		hashedPassword, err := s.HashPassword(req.Password)
		if err != nil {
			return err
		}
		student.Password = hashedPassword
	}

	return s.dao.UpdateStudent(&student)
}

// DeleteStudent 删除学生
func (s *Service) DeleteStudent(id uint) error {
	return s.dao.DeleteStudent(id)
}

// GetTasks 获取任务列表
func (s *Service) GetTasks() (tasks []models.Task, err error) {
	return s.dao.GetTasks()
}

// CreateTask 创建任务
func (s *Service) CreateTask(req *models.CreateTaskRequest, teacherID uint) error {
	task := &models.Task{
		Name:           req.Name,
		Description:    req.Description,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
		CompletedCount: 0,
		Status:         "pending",
		TeacherID:      teacherID,
	}

	// 创建任务
	err := s.dao.CreateTask(task)
	if err != nil {
		return err
	}

	// 如果指定了学生，则建立关联关系
	if len(req.StudentIDs) > 0 {
		students, err := s.dao.GetStudentsByIDs(req.StudentIDs)
		if err != nil {
			return err
		}
		task.Students = students
		return s.dao.UpdateTask(task)
	}

	return nil
}

// UpdateTask 更新任务
func (s *Service) UpdateTask(id uint, req *models.UpdateTaskRequest) error {
	task, err := s.dao.GetTaskByID(id)
	if err != nil {
		return err
	}

	task.Name = req.Name
	task.Description = req.Description
	task.StartTime = req.StartTime
	task.EndTime = req.EndTime

	// 如果指定了学生，则更新关联关系
	if len(req.StudentIDs) > 0 {
		students, err := s.dao.GetStudentsByIDs(req.StudentIDs)
		if err != nil {
			return err
		}
		task.Students = students
	}

	return s.dao.UpdateTask(&task)
}

// UpdateTaskStatus 更新任务状态
func (s *Service) UpdateTaskStatus(id uint, status string) error {
	return s.dao.UpdateTaskStatus(id, status)
}

// DeleteTask 删除任务
func (s *Service) DeleteTask(id uint) error {
	return s.dao.DeleteTask(id)
}

// GetTaskStudents 获取任务的学生完成情况
func (s *Service) GetTaskStudents(taskID uint) (taskStudents []models.TaskStudent, err error) {
	return s.dao.GetTaskStudents(taskID)
}

// UpdateStudentTaskStatus 更新学生任务完成状态
func (s *Service) UpdateStudentTaskStatus(taskID, studentID uint, completed bool) error {
	// 更新学生任务完成状态
	err := s.dao.UpdateStudentTaskStatus(taskID, studentID, completed)
	if err != nil {
		return err
	}

	// 更新任务的完成数量
	completedCount, _, err := s.dao.GetTaskCompletionStats(taskID)
	if err != nil {
		return err
	}

	// 更新任务的completed_count字段
	return s.dao.UpdateTaskCompletedCount(taskID, completedCount)
}

// GetTaskCompletionStats 获取任务完成统计
func (s *Service) GetTaskCompletionStats(taskID uint) (completedCount, totalCount int, err error) {
	return s.dao.GetTaskCompletionStats(taskID)
}
