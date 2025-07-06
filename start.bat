@echo off
echo 启动任务管理系统...

echo 启动后端服务器...
cd backend
start "Backend Server" cmd /k "go mod tidy && go run main.go"

echo 等待后端启动...
timeout /t 3 /nobreak > nul

echo 启动前端开发服务器...
cd ../frontend
start "Frontend Server" cmd /k "npm install && npm run dev"

echo 系统启动完成！
echo 前端地址: http://localhost:3000
echo 后端地址: http://localhost:8080
echo.
echo 默认账户:
echo 管理员 - admin / admin123
echo 学生 - student / student123
pause 