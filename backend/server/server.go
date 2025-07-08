package server

import (
	"backend/logger"
	"backend/middleware"
	"backend/pkg/settings"
	"backend/service"

	"github.com/gin-gonic/gin"
)

var svc *service.Service

func initRouter() *gin.Engine {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 注册使用的中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.CORSMiddleware())

	// 后台
	backend := r.Group("/backend")
	backend.POST("/login", login)

	// 前台API
	frontend := r.Group("/api")
	{
		// 学生管理接口
		frontend.GET("/students", getStudents)
		frontend.POST("/students", createStudent)
		frontend.PUT("/students/:id", updateStudent)
		frontend.DELETE("/students/:id", deleteStudent)

		// 任务管理接口
		frontend.GET("/tasks", getTasks)
		frontend.POST("/tasks", createTask)
		frontend.PUT("/tasks/:id", updateTask)
		frontend.PUT("/tasks/:id/status", updateTaskStatus)
		frontend.DELETE("/tasks/:id", deleteTask)

		// 任务学生完成状态接口
		frontend.GET("/tasks/:id/students", getTaskStudents)
		frontend.PUT("/tasks/:id/students/:studentId/status", updateStudentTaskStatus)
		frontend.GET("/tasks/:id/completion-stats", getTaskCompletionStats)
	}

	return r
}

func Init(app *settings.AppConfig) *gin.Engine {
	svc = service.InitService(app)
	return initRouter()
}
