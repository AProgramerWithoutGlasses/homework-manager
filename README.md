# 作业管理系统

这是一个用于管理作业和学生信息的系统，包含前端和后端两个部分。

## 项目结构

```
homework-manager/
├── backend/          # 后端代码 (Go + Gin)
├── frontend/         # 前端代码 (Vue 3 + Element Plus)
├── README.md         # 项目说明
└── start.bat         # 启动脚本
```

## 功能特性

- 老师登录认证
- 学生信息管理（增删改查）
- 任务/作业管理（增删改查）
- 任务状态跟踪
- 响应式界面设计

## 技术栈

### 后端

- Go 1.21+
- Gin Web 框架
- GORM ORM 框架
- MySQL 数据库
- Redis 缓存
- JWT 认证

### 前端

- Vue 3
- Element Plus UI 组件库
- Pinia 状态管理
- Vue Router 路由管理
- Axios HTTP 客户端

## 快速开始

### 1. 环境准备

确保已安装以下软件：

- Go 1.21+
- Node.js 16+
- MySQL 8.0+
- Redis 6.0+

### 2. 数据库配置

1. 创建 MySQL 数据库：

```sql
CREATE DATABASE shoushou CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 修改后端配置文件 `backend/configs/local.yaml`：

```yaml
mysql:
  host: "localhost"
  port: 3306
  user: "root"
  password: "你的密码"
  dbname: "shoushou"
```

### 3. 启动后端

```bash
cd backend
go mod tidy
go run main.go
```

后端将在 `http://localhost:8080` 启动

### 4. 启动前端

```bash
cd frontend
npm install
npm run dev
```

前端将在 `http://localhost:3000` 启动

### 5. 使用启动脚本

也可以直接运行项目根目录的 `start.bat` 脚本来同时启动前后端。

### 6. 默认账户

系统会自动创建默认管理员账户：

| 电话号码 | 密码     | 角色   |
| -------- | -------- | ------ |
| admin    | admin123 | 管理员 |

## API 接口文档

### 认证接口

#### 老师登录

- **URL**: `POST /backend/login`
- **请求体**:

```json
{
  "telephone": "admin",
  "password": "admin123"
}
```

- **响应**:

```json
{
  "code": 200,
  "data": {
    "token": "jwt_token",
    "teacher": {
      "id": 1,
      "telephone": "admin",
      "name": "管理员"
    }
  }
}
```

### 学生管理接口

#### 获取学生列表

- **URL**: `GET /api/students`
- **响应**:

```json
{
  "code": 200,
  "data": {
    "students": [
      {
        "id": 1,
        "username": "zhangsan",
        "name": "张三",
        "class": "计算机1班",
        "grade": "大一"
      }
    ]
  }
}
```

#### 创建学生

- **URL**: `POST /api/students`
- **请求体**:

```json
{
  "username": "lisi",
  "password": "123456",
  "name": "李四",
  "class": "计算机1班",
  "grade": "大一"
}
```

#### 更新学生信息

- **URL**: `PUT /api/students/:id`
- **请求体**:

```json
{
  "username": "lisi",
  "password": "newpassword",
  "name": "李四",
  "class": "计算机1班",
  "grade": "大一"
}
```

#### 删除学生

- **URL**: `DELETE /api/students/:id`

### 任务管理接口

#### 获取任务列表

- **URL**: `GET /api/tasks`
- **响应**:

```json
{
  "code": 200,
  "data": {
    "tasks": [
      {
        "id": 1,
        "name": "完成作业1",
        "start_time": "2024-01-15T09:00:00.000Z",
        "end_time": "2024-01-20T18:00:00.000Z",
        "completed_count": 5,
        "total_count": 10,
        "status": "pending",
        "teacher_id": 1,
        "teacher": {
          "id": 1,
          "telephone": "admin",
          "name": "管理员"
        },
        "students": [
          {
            "id": 1,
            "username": "zhangsan",
            "name": "张三",
            "class": "计算机1班",
            "grade": "大一"
          }
        ]
      }
    ]
  }
}
```

#### 创建任务

- **URL**: `POST /api/tasks`
- **请求体**:

```json
{
  "name": "新任务",
  "start_time": "2024-01-15T09:00:00.000Z",
  "end_time": "2024-01-20T18:00:00.000Z",
  "total_count": 10,
  "student_ids": [1, 2, 3]
}
```

#### 更新任务

- **URL**: `PUT /api/tasks/:id`
- **请求体**: 同创建任务

#### 更新任务状态

- **URL**: `PUT /api/tasks/:id/status`
- **请求体**:

```json
{
  "status": "completed"
}
```

#### 删除任务

- **URL**: `DELETE /api/tasks/:id`

#### 获取任务学生完成情况

- **URL**: `GET /api/tasks/:id/students`
- **响应**:

```json
{
  "code": 200,
  "data": {
    "task_students": [
      {
        "task_id": 1,
        "student_id": 1,
        "completed": true,
        "completed_at": "2024-01-15T10:30:00Z"
      }
    ]
  }
}
```

#### 更新学生任务完成状态

- **URL**: `PUT /api/tasks/:id/students/:studentId/status`
- **请求体**:

```json
{
  "completed": true
}
```

#### 获取任务完成统计

- **URL**: `GET /api/tasks/:id/completion-stats`
- **响应**:

```json
{
  "code": 200,
  "data": {
    "completed_count": 5,
    "total_count": 10
  }
}
```

## 状态码说明

- `200`: 成功
- `400`: 参数错误
- `401`: 未授权
- `500`: 服务器错误

## 开发说明

### 后端开发

- 使用 Gin 框架处理 HTTP 请求
- 使用 GORM 进行数据库操作
- 使用 JWT 进行用户认证
- 使用 Zap 进行日志记录

### 前端开发

- 使用 Vue 3 Composition API
- 使用 Element Plus 组件库
- 使用 Pinia 进行状态管理
- 使用 Axios HTTP 客户端

## 许可证

MIT License
