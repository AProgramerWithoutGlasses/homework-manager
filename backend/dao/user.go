package dao

import (
	"backend/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

// GetTeacherByTelephone 根据电话号码获取老师
func (d *Dao) GetTeacherByTelephone(telephone string) (teacher models.Teacher, err error) {
	err = d.db.Where("telephone = ?", telephone).First(&teacher).Error
	return
}

// GetStudents 获取学生列表
func (d *Dao) GetStudents() (students []models.Student, err error) {
	err = d.db.Find(&students).Error
	return
}

// CreateStudent 创建学生
func (d *Dao) CreateStudent(student *models.Student) error {
	return d.db.Create(student).Error
}

// GetStudentByID 根据ID获取学生
func (d *Dao) GetStudentByID(id uint) (student models.Student, err error) {
	err = d.db.First(&student, id).Error
	return
}

// UpdateStudent 更新学生信息
func (d *Dao) UpdateStudent(student *models.Student) error {
	return d.db.Save(student).Error
}

// DeleteStudent 删除学生
func (d *Dao) DeleteStudent(id uint) error {
	return d.db.Delete(&models.Student{}, id).Error
}

// GetTasks 获取任务列表
func (d *Dao) GetTasks() (tasks []models.Task, err error) {
	err = d.db.Preload("Teacher").Preload("Students").Find(&tasks).Error
	return
}

// CreateTask 创建任务
func (d *Dao) CreateTask(task *models.Task) error {
	return d.db.Create(task).Error
}

// GetTaskByID 根据ID获取任务
func (d *Dao) GetTaskByID(id uint) (task models.Task, err error) {
	err = d.db.Preload("Teacher").Preload("Students").First(&task, id).Error
	return
}

// UpdateTask 更新任务
func (d *Dao) UpdateTask(task *models.Task) error {
	return d.db.Save(task).Error
}

// UpdateTaskStatus 更新任务状态
func (d *Dao) UpdateTaskStatus(id uint, status string) error {
	return d.db.Model(&models.Task{}).Where("id = ?", id).Update("status", status).Error
}

// DeleteTask 删除任务
func (d *Dao) DeleteTask(id uint) error {
	return d.db.Delete(&models.Task{}, id).Error
}

// GetStudentsByIDs 根据ID列表获取学生
func (d *Dao) GetStudentsByIDs(ids []uint) (students []models.Student, err error) {
	err = d.db.Where("id IN ?", ids).Find(&students).Error
	return
}

// GetTaskStudents 获取任务的学生完成情况
func (d *Dao) GetTaskStudents(taskID uint) (taskStudents []models.TaskStudent, err error) {
	err = d.db.Where("task_id = ?", taskID).Find(&taskStudents).Error
	return
}

// UpdateStudentTaskStatus 更新学生任务完成状态
func (d *Dao) UpdateStudentTaskStatus(taskID, studentID uint, completed bool) error {
	var taskStudent models.TaskStudent
	err := d.db.Where("task_id = ? AND student_id = ?", taskID, studentID).First(&taskStudent).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果记录不存在，创建新记录
			taskStudent = models.TaskStudent{
				TaskID:    taskID,
				StudentID: studentID,
				Completed: completed,
			}
			if completed {
				now := time.Now()
				taskStudent.CompletedAt = &now
			}
			return d.db.Create(&taskStudent).Error
		}
		return err
	}

	// 更新完成状态
	taskStudent.Completed = completed
	if completed && taskStudent.CompletedAt == nil {
		now := time.Now()
		taskStudent.CompletedAt = &now
	} else if !completed {
		taskStudent.CompletedAt = nil
	}

	return d.db.Save(&taskStudent).Error
}

// GetTaskCompletionStats 获取任务完成统计
func (d *Dao) GetTaskCompletionStats(taskID uint) (completedCount, totalCount int, err error) {
	var completedCount64, totalCount64 int64

	// 获取总学生数
	err = d.db.Model(&models.TaskStudent{}).Where("task_id = ?", taskID).Count(&totalCount64).Error
	if err != nil {
		return
	}

	// 获取已完成学生数
	err = d.db.Model(&models.TaskStudent{}).Where("task_id = ? AND completed = ?", taskID, true).Count(&completedCount64).Error
	if err != nil {
		return
	}

	completedCount = int(completedCount64)
	totalCount = int(totalCount64)
	return
}

// UpdateTaskCompletedCount 更新任务的完成数量
func (d *Dao) UpdateTaskCompletedCount(taskID uint, completedCount int) error {
	return d.db.Model(&models.Task{}).Where("id = ?", taskID).Update("completed_count", completedCount).Error
}
