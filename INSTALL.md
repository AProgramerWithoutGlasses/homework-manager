# 安装指南

## 系统要求

### 后端要求

- Go 1.21 或更高版本
- Git

### 前端要求

- Node.js 16 或更高版本
- npm 或 yarn

## 安装步骤

### 1. 安装 Go

访问 [Go 官网](https://golang.org/dl/) 下载并安装 Go。

验证安装：

```bash
go version
```

### 2. 安装 Node.js

访问 [Node.js 官网](https://nodejs.org/) 下载并安装 Node.js。

验证安装：

```bash
node --version
npm --version
```

### 3. 克隆项目

```bash
git clone <repository-url>
cd task-management-system
```

### 4. 安装后端依赖

```bash
cd backend
go mod tidy
```

### 5. 安装前端依赖

```bash
cd frontend
npm install
```

## 启动系统

### 方法一：使用启动脚本（推荐）

双击运行 `start.bat` 文件（Windows）

### 方法二：手动启动

#### 启动后端

```bash
cd backend
go run main.go
```

#### 启动前端

```bash
cd frontend
npm run dev
```

## 访问系统

- 前端地址：http://localhost:3000
- 后端 API：http://localhost:8080

## 默认账户

- 管理员：admin / admin123
- 学生：student / student123

## 故障排除

### 后端启动失败

1. 检查 Go 版本是否满足要求
2. 确保在 backend 目录下运行命令
3. 检查端口 8080 是否被占用

### 前端启动失败

1. 检查 Node.js 版本是否满足要求
2. 确保在 frontend 目录下运行命令
3. 检查端口 3000 是否被占用

### 数据库问题

1. 确保有写入权限
2. 删除 `backend/task_management.db` 文件重新创建

## 开发模式

### 后端热重载

```bash
go install github.com/cosmtrek/air@latest
air
```

### 前端热重载

前端已配置热重载，修改代码后会自动刷新。
