package server

import (
	"backend/models"
	"backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// login 登录接口
func login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("login", zap.Error(err))
		response.Fail(c, response.ParamErrCode)
		return
	}

	// 调用service层登录方法
	loginResp, err := svc.Login(&req)
	if err != nil {
		zap.L().Error("login", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, loginResp)
}

// getStudents 获取学生列表
func getStudents(c *gin.Context) {
	students, err := svc.GetStudents()
	if err != nil {
		zap.L().Error("getStudents", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, models.StudentsResponse{Students: students})
}

// createStudent 创建学生
func createStudent(c *gin.Context) {
	var req models.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("createStudent", zap.Error(err))
		response.Fail(c, response.ParamErrCode)
		return
	}

	err := svc.CreateStudent(&req)
	if err != nil {
		zap.L().Error("createStudent", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, nil)
}

// updateStudent 更新学生信息
func updateStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Fail(c, response.ParamErrCode)
		return
	}

	var req models.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("updateStudent", zap.Error(err))
		response.Fail(c, response.ParamErrCode)
		return
	}

	err = svc.UpdateStudent(uint(id), &req)
	if err != nil {
		zap.L().Error("updateStudent", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, nil)
}

// deleteStudent 删除学生
func deleteStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Fail(c, response.ParamErrCode)
		return
	}

	err = svc.DeleteStudent(uint(id))
	if err != nil {
		zap.L().Error("deleteStudent", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, nil)
}

// getTasks 获取任务列表
func getTasks(c *gin.Context) {
	tasks, err := svc.GetTasks()
	if err != nil {
		zap.L().Error("getTasks", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, models.TasksResponse{Tasks: tasks})
}

// createTask 创建任务
func createTask(c *gin.Context) {
	var req models.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("createTask", zap.Error(err))
		response.Fail(c, response.ParamErrCode)
		return
	}

	// 从JWT中获取老师ID，这里暂时使用固定值1
	teacherID := uint(1)
	err := svc.CreateTask(&req, teacherID)
	if err != nil {
		zap.L().Error("createTask", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, nil)
}

// updateTask 更新任务
func updateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Fail(c, response.ParamErrCode)
		return
	}

	var req models.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("updateTask", zap.Error(err))
		response.Fail(c, response.ParamErrCode)
		return
	}

	err = svc.UpdateTask(uint(id), &req)
	if err != nil {
		zap.L().Error("updateTask", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, nil)
}

// updateTaskStatus 更新任务状态
func updateTaskStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Fail(c, response.ParamErrCode)
		return
	}

	var req models.UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("updateTaskStatus", zap.Error(err))
		response.Fail(c, response.ParamErrCode)
		return
	}

	err = svc.UpdateTaskStatus(uint(id), req.Status)
	if err != nil {
		zap.L().Error("updateTaskStatus", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, nil)
}

// deleteTask 删除任务
func deleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Fail(c, response.ParamErrCode)
		return
	}

	err = svc.DeleteTask(uint(id))
	if err != nil {
		zap.L().Error("deleteTask", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, nil)
}

// getTaskStudents 获取任务的学生完成情况
func getTaskStudents(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Fail(c, response.ParamErrCode)
		return
	}

	taskStudents, err := svc.GetTaskStudents(uint(id))
	if err != nil {
		zap.L().Error("getTaskStudents", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, models.TaskStudentsResponse{TaskStudents: taskStudents})
}

// updateStudentTaskStatus 更新学生任务完成状态
func updateStudentTaskStatus(c *gin.Context) {
	taskIDStr := c.Param("id")
	studentIDStr := c.Param("studentId")

	taskID, err := strconv.ParseUint(taskIDStr, 10, 32)
	if err != nil {
		response.Fail(c, response.ParamErrCode)
		return
	}

	studentID, err := strconv.ParseUint(studentIDStr, 10, 32)
	if err != nil {
		response.Fail(c, response.ParamErrCode)
		return
	}

	var req models.UpdateStudentTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("updateStudentTaskStatus", zap.Error(err))
		response.Fail(c, response.ParamErrCode)
		return
	}

	err = svc.UpdateStudentTaskStatus(uint(taskID), uint(studentID), req.Completed)
	if err != nil {
		zap.L().Error("updateStudentTaskStatus", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, nil)
}

// getTaskCompletionStats 获取任务完成统计
func getTaskCompletionStats(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Fail(c, response.ParamErrCode)
		return
	}

	completedCount, totalCount, err := svc.GetTaskCompletionStats(uint(id))
	if err != nil {
		zap.L().Error("getTaskCompletionStats", zap.Error(err))
		response.FailWithMsg(c, response.ServerErrCode, err.Error())
		return
	}

	response.Success(c, models.TaskCompletionStatsResponse{
		CompletedCount: completedCount,
		TotalCount:     totalCount,
	})
}
