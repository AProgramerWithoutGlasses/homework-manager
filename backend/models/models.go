package models

import (
	"time"

	"gorm.io/gorm"
)

// Teacher 老师表
type Teacher struct {
	gorm.Model
	Telephone string `gorm:"type:varchar(50);not null;uniqueIndex:idx_telephone_deleted_at" json:"telephone"`
	Password  string `gorm:"type:varchar(100);not null" json:"-"` // 密码不返回给前端
	Name      string `gorm:"type:varchar(50);not null" json:"name"`

	// 关联关系
	Tasks []Task `gorm:"foreignKey:TeacherID" json:"tasks"`
}

// Student 学生表
type Student struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;uniqueIndex:idx_username_deleted_at" json:"username"`
	Password string `gorm:"type:varchar(100);not null" json:"-"` // 密码不返回给前端
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	Class    string `gorm:"type:varchar(50)" json:"class"`
	Grade    string `gorm:"type:varchar(20)" json:"grade"`

	// 关联关系
	Tasks []Task `gorm:"many2many:task_students;" json:"tasks"`
}

// Task 任务表
type Task struct {
	gorm.Model
	Name           string    `gorm:"type:varchar(200);not null" json:"name"`
	Description    string    `gorm:"type:varchar(500)" json:"description"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	CompletedCount int       `gorm:"default:0" json:"completed_count"` // 已完成该任务的学生数量
	Status         string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	TeacherID      uint      `gorm:"not null" json:"teacher_id"` // 老师ID

	// 关联关系
	Teacher  *Teacher  `gorm:"foreignKey:TeacherID" json:"teacher"`
	Students []Student `gorm:"many2many:task_students;" json:"students"`
}

// TaskStudent 任务学生关联表
type TaskStudent struct {
	TaskID      uint       `gorm:"primaryKey"`
	StudentID   uint       `gorm:"primaryKey"`
	Completed   bool       `gorm:"default:false" json:"completed"` // 学生是否完成该任务
	CompletedAt *time.Time `json:"completed_at"`                   // 完成时间
}
