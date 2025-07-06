package models

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;uniqueIndex:idx_username_deleted_at" json:"username"`
	Password string `gorm:"type:varchar(100);not null" json:"-"` // 密码不返回给前端
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	Class    string `gorm:"type:varchar(50)" json:"class"`
}
