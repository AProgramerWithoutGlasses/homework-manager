<template>
  <div class="dashboard">
    <el-container>
      <el-aside width="250px">
        <div class="sidebar">
          <div class="logo">
            <h2>任务管理系统</h2>
          </div>
          
          <el-menu
            :default-active="$route.path"
            router
            class="sidebar-menu"
          >
            <el-menu-item index="/dashboard/students">
              <el-icon><User /></el-icon>
              <span>学生管理</span>
            </el-menu-item>
            <el-menu-item index="/dashboard/tasks">
              <el-icon><List /></el-icon>
              <span>任务管理</span>
            </el-menu-item>
          </el-menu>
        </div>
      </el-aside>
      
      <el-container>
        <el-header>
          <div class="header-content">
            <div class="header-left">
              <h3>{{ getPageTitle() }}</h3>
            </div>
            <div class="header-right">
              <el-dropdown @command="handleCommand">
                <span class="user-info">
                  {{ authStore.teacher?.name || '用户' }}
                  <el-icon><ArrowDown /></el-icon>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item command="logout">退出登录</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
          </div>
        </el-header>
        
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

export default {
  name: 'Dashboard',
  setup() {
    const authStore = useAuthStore()
    const router = useRouter()
    
    const getPageTitle = () => {
      const path = router.currentRoute.value.path
      if (path.includes('students')) return '学生管理'
      if (path.includes('tasks')) return '任务管理'
      return '仪表板'
    }
    
    const handleCommand = (command) => {
      if (command === 'logout') {
        authStore.logout()
        router.push('/login')
        ElMessage.success('已退出登录')
      }
    }
    
    return {
      authStore,
      getPageTitle,
      handleCommand
    }
  }
}
</script>

<style scoped>
.dashboard,
.el-container {
  height: 100vh;
  min-height: 100vh;
}

.el-aside {
  height: 100vh;
  background: #304156;
  padding: 0;
  border: none;
}

.sidebar {
  height: 100%;
  background: #304156;
  color: white;
  display: flex;
  flex-direction: column;
}

.logo {
  padding: 20px;
  text-align: center;
  border-bottom: 1px solid #435266;
}

.logo h2 {
  color: white;
  margin: 0;
  font-size: 18px;
}

.sidebar-menu {
  border: none;
  background: transparent;
}

.el-header {
  background: white;
  border-bottom: 1px solid #e6e6e6;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
}

.header-content {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-left h3 {
  margin: 0;
  color: #333;
}

.user-info {
  display: flex;
  align-items: center;
  cursor: pointer;
  color: #666;
}

.user-info:hover {
  color: #409eff;
}

.el-main {
  background: #f5f5f5;
  padding: 20px;
}
</style> 