package dao

import (
	"backend/models"
	"backend/pkg/settings"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Dao struct {
	db  *gorm.DB
	rdb *redis.Client
}

func Init(app *settings.AppConfig) *Dao {
	db := initDB(app.MySQLConfig)

	// 自动迁移数据库表
	err := db.AutoMigrate(&models.Teacher{}, &models.Student{}, &models.Task{}, &models.TaskStudent{})
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}

	// 初始化测试数据
	initTestData(db)

	dao := &Dao{
		db:  db,
		rdb: initRDB(app.RedisConfig),
	}
	return dao
}

// initTestData 初始化测试数据
func initTestData(db *gorm.DB) {
	// 检查是否已有老师数据
	var teacherCount int64
	db.Model(&models.Teacher{}).Count(&teacherCount)
	if teacherCount == 0 {
		// 创建默认管理员老师
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		adminTeacher := models.Teacher{
			Telephone: "admin",
			Password:  string(hashedPassword),
			Name:      "管理员",
		}
		db.Create(&adminTeacher)
	}

	// 检查是否已有学生数据
	var studentCount int64
	db.Model(&models.Student{}).Count(&studentCount)
	if studentCount == 0 {
		// 创建测试学生数据
		students := []models.Student{
			{Username: "zhangsan", Password: "123456", Name: "张三", Class: "计算机1班", Grade: "大一"},
			{Username: "lisi", Password: "123456", Name: "李四", Class: "计算机1班", Grade: "大一"},
			{Username: "wangwu", Password: "123456", Name: "王五", Class: "计算机2班", Grade: "大一"},
		}

		// 批量生成100条学生数据
		classList := []string{"计算机1班", "计算机2班", "软件1班", "软件2班", "网络1班"}
		gradeList := []string{"大一", "大二", "大三", "大四"}
		for i := 1; i <= 100; i++ {
			username := fmt.Sprintf("student%03d", i)
			name := fmt.Sprintf("学生%d", i)
			className := classList[(i-1)%len(classList)]
			gradeName := gradeList[(i-1)%len(gradeList)]
			students = append(students, models.Student{
				Username: username,
				Password: "123456",
				Name:     name,
				Class:    className,
				Grade:    gradeName,
			})
		}

		for _, student := range students {
			// 加密密码
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
			student.Password = string(hashedPassword)
			db.Create(&student)
		}
	}

	// 检查是否已有任务数据
	var taskCount int64
	db.Model(&models.Task{}).Count(&taskCount)
	if taskCount == 0 {
		// 创建测试任务数据
		startTime1, _ := time.Parse("2006-01-02 15:04:05", "2024-01-15 09:00:00")
		endTime1, _ := time.Parse("2006-01-02 15:04:05", "2024-01-20 18:00:00")
		startTime2, _ := time.Parse("2006-01-02 15:04:05", "2024-01-16 10:00:00")
		endTime2, _ := time.Parse("2006-01-02 15:04:05", "2024-01-25 18:00:00")
		startTime3, _ := time.Parse("2006-01-02 15:04:05", "2024-01-18 14:00:00")
		endTime3, _ := time.Parse("2006-01-02 15:04:05", "2024-01-30 18:00:00")

		tasks := []models.Task{
			{Name: "完成作业1", Description: "完成作业1的相关内容", StartTime: startTime1, EndTime: endTime1, CompletedCount: 5, Status: "pending", TeacherID: 1},
			{Name: "复习课程", Description: "复习课程的相关内容", StartTime: startTime2, EndTime: endTime2, CompletedCount: 8, Status: "in_progress", TeacherID: 1},
			{Name: "准备考试", Description: "准备考试的相关内容", StartTime: startTime3, EndTime: endTime3, CompletedCount: 15, Status: "completed", TeacherID: 1},
		}

		for _, task := range tasks {
			db.Create(&task)
		}
	}
}
