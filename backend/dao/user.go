package dao

import (
	"backend/models"
)

// GetUserByUsername 根据用户名获取用户信息
func (d *Dao) GetUserByUsername(username string) (user models.User, err error) {
	err = d.db.Preload("Company").Preload("Periods").Where("username = ?", username).First(&user).Error
	return
}
