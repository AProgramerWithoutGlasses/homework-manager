# 任务管理系统

这是一个基于 Vue3 前端和 Go 后端的完整任务管理系统，提供用户登录、学生管理和任务管理功能。

## 🚀 功能特性

### 用户认证

- 用户登录/退出
- JWT 令牌认证
- 角色权限管理（管理员/学生）

### 学生管理

- 学生信息的增删改查
- 学生状态管理（活跃/非活跃）
- 学生信息包括：姓名、邮箱、电话、班级、年级

### 任务管理

- 任务的增删改查
- 任务状态跟踪（待处理、进行中、已完成、已取消）
- 任务优先级管理（低、中、高）
- 任务分配给指定学生
- 截止日期设置

## 🛠️ 技术栈

### 后端 (Go)

- **框架**: Gin
- **ORM**: GORM
- **数据库**: SQLite
- **认证**: JWT
- **密码加密**: bcrypt

### 前端 (Vue3)

- **框架**: Vue 3 (Composition API)
- **路由**: Vue Router 4
- **状态管理**: Pinia
- **UI 组件**: Element Plus
- **HTTP 客户端**: Axios
- **构建工具**: Vite

## 📁 项目结构

```
task-management-system/
├── backend/                 # Go后端
│   ├── config/             # 数据库配置
│   ├── controllers/        # 控制器
│   ├── middleware/         # 中间件
│   ├── models/            # 数据模型
│   └── main.go           # 主程序
├── frontend/              # Vue3前端
│   ├── src/
│   │   ├── views/        # 页面组件
│   │   ├── stores/       # 状态管理
│   │   └── router/       # 路由配置
│   ├── package.json
│   └── vite.config.js
├── start.bat             # Windows启动脚本
└── README.md
```

## 🚀 快速开始

### 环境要求

- Go 1.21+
- Node.js 16+
- npm 或 yarn

### 1. 克隆项目

```bash
git clone <repository-url>
cd task-management-system
```

### 2. 启动后端

```bash
cd backend
go mod tidy
go run main.go
```

后端将在 http://localhost:8080 启动

### 3. 启动前端

```bash
cd frontend
npm install
npm run dev
```

前端将在 http://localhost:3000 启动

### 4. 使用启动脚本（Windows）

双击运行 `start.bat` 文件，将自动启动前后端服务。

## 👤 默认账户

系统预置了两个默认账户：

| 角色   | 用户名  | 密码       |
| ------ | ------- | ---------- |
| 管理员 | admin   | admin123   |
| 学生   | student | student123 |

## 📋 API 接口

### 认证接口

- `POST /api/login` - 用户登录
- `POST /api/register` - 用户注册

### 学生管理接口

- `GET /api/students` - 获取学生列表
- `POST /api/students` - 创建学生
- `GET /api/students/:id` - 获取学生详情
- `PUT /api/students/:id` - 更新学生信息
- `DELETE /api/students/:id` - 删除学生

### 任务管理接口

- `GET /api/tasks` - 获取任务列表
- `POST /api/tasks` - 创建任务
- `GET /api/tasks/:id` - 获取任务详情
- `PUT /api/tasks/:id` - 更新任务
- `DELETE /api/tasks/:id` - 删除任务
- `PUT /api/tasks/:id/status` - 更新任务状态

## 🎨 界面预览

### 登录页面

- 现代化的登录界面
- 表单验证
- 错误提示

### 主界面

- 响应式侧边栏导航
- 用户信息显示
- 退出登录功能

### 学生管理

- 学生列表表格
- 添加/编辑学生对话框
- 状态标签显示
- 删除确认

### 任务管理

- 任务列表表格
- 状态和优先级标签
- 添加/编辑任务对话框
- 状态更新功能

## 🔧 开发说明

### 后端开发

1. 修改数据库配置：`backend/config/database.go`
2. 添加新的模型：`backend/models/`
3. 创建新的控制器：`backend/controllers/`
4. 更新路由：`backend/main.go`

### 前端开发

1. 添加新页面：`frontend/src/views/`
2. 更新路由：`frontend/src/router/index.js`
3. 添加状态管理：`frontend/src/stores/`

## 📝 注意事项

1. 首次运行时会自动创建 SQLite 数据库文件
2. 系统会自动创建默认用户账户
3. 前端开发服务器配置了 API 代理，无需额外配置
4. 密码使用 bcrypt 加密存储

## 🤝 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目！

## �� 许可证

MIT License
