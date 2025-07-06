package models

// LoginResponse 登录响应结构体
type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
