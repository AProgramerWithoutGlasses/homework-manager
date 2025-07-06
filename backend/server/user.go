package server

import (
	"backend/models"
	"backend/pkg/response"
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
