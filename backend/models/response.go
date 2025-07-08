package models

// LoginResponse 登录响应结构体
type LoginResponse struct {
	Token   string   `json:"token"`
	Teacher *Teacher `json:"teacher"`
}

// StudentsResponse 学生列表响应结构体
type StudentsResponse struct {
	Students []Student `json:"students"`
}

// TasksResponse 任务列表响应结构体
type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}

// TaskStudentsResponse 任务学生完成情况响应结构体
type TaskStudentsResponse struct {
	TaskStudents []TaskStudent `json:"task_students"`
}

// TaskCompletionStatsResponse 任务完成统计响应结构体
type TaskCompletionStatsResponse struct {
	CompletedCount int `json:"completed_count"`
	TotalCount     int `json:"total_count"`
}
