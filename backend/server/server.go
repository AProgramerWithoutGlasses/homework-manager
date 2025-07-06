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

	// 前台
	//frontend := r.Group("/frontend")

	return r
}

func Init(app *settings.AppConfig) *gin.Engine {
	svc = service.InitService(app)
	return initRouter()
}
