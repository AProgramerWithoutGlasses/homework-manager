package models

import "time"

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Telephone string `json:"telephone" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// CreateStudentRequest 创建学生请求结构体
type CreateStudentRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Class    string `json:"class"`
	Grade    string `json:"grade"`
}

// UpdateStudentRequest 更新学生请求结构体
type UpdateStudentRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Name     string `json:"name" binding:"required"`
	Class    string `json:"class"`
	Grade    string `json:"grade"`
}

// CreateTaskRequest 创建任务请求结构体
type CreateTaskRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	StudentIDs  []uint    `json:"student_ids"`
}

// UpdateTaskRequest 更新任务请求结构体
type UpdateTaskRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	StudentIDs  []uint    `json:"student_ids"`
}

// UpdateTaskStatusRequest 更新任务状态请求结构体
type UpdateTaskStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

// UpdateStudentTaskStatusRequest 更新学生任务完成状态请求结构体
type UpdateStudentTaskStatusRequest struct {
	Completed bool `json:"completed" binding:"required"`
}
