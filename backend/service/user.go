package service

import (
	"backend/middleware"
	"backend/models"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login 用户登录
func (s *Service) Login(req *models.LoginRequest) (res *models.LoginResponse, err error) {
	// 获取用户
	user, err := s.dao.GetUserByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("账户不存在")
		}
		return nil, fmt.Errorf("s.dao.GetUserByUsername() err: %v", err)
	}

	// 验证密码
	if !s.CheckPassword(user.Password, req.Password) {
		return nil, errors.New("密码错误")
	}

	// 生成JWT令牌
	token, err := middleware.GenToken(user.ID, user.Username)
	if err != nil {
		return nil, fmt.Errorf("middleware.GenToken() err: %v", err)
	}

	res = &models.LoginResponse{
		Token: token,
		User:  &user,
	}

	return
}

// CheckPassword 验证密码
func (s *Service) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
