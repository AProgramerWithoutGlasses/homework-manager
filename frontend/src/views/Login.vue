<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-header">
        <h1>任务管理系统</h1>
        <p>请登录您的账户</p>
      </div>
      
      <el-form 
        ref="loginForm" 
        :model="loginData" 
        :rules="rules" 
        class="login-form"
        @submit.prevent="handleLogin"
      >
        <el-form-item prop="username">
          <el-input
            v-model="loginData.username"
            placeholder="用户名"
            prefix-icon="User"
            size="large"
          />
        </el-form-item>
        
        <el-form-item prop="password">
          <el-input
            v-model="loginData.password"
            type="password"
            placeholder="密码"
            prefix-icon="Lock"
            size="large"
            show-password
          />
        </el-form-item>
        
        <el-form-item>
          <el-button
            type="primary"
            size="large"
            :loading="loading"
            @click="handleLogin"
            style="width: 100%"
          >
            登录
          </el-button>
        </el-form-item>
      </el-form>
      
      <div class="login-tips">
        <p>默认账户：</p>
        <p>管理员 - admin / admin123</p>
        <p>学生 - student / student123</p>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { ElMessage } from 'element-plus'

export default {
  name: 'Login',
  setup() {
    const router = useRouter()
    const authStore = useAuthStore()
    const loginForm = ref(null)
    const loading = ref(false)
    
    const loginData = reactive({
      username: '',
      password: ''
    })
    
    const rules = {
      username: [
        { required: true, message: '请输入用户名', trigger: 'blur' }
      ],
      password: [
        { required: true, message: '请输入密码', trigger: 'blur' }
      ]
    }
    
    const handleLogin = async () => {
      if (!loginForm.value) return
      
      try {
        await loginForm.value.validate()
        loading.value = true
        
        const result = await authStore.login(loginData.username, loginData.password)
        
        if (result.success) {
          ElMessage.success('登录成功')
          router.push('/dashboard')
        } else {
          ElMessage.error(result.error)
        }
      } catch (error) {
        ElMessage.error('登录失败')
      } finally {
        loading.value = false
      }
    }
    
    return {
      loginForm,
      loginData,
      rules,
      loading,
      handleLogin
    }
  }
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-box {
  background: white;
  border-radius: 12px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  padding: 40px;
  width: 100%;
  max-width: 400px;
}

.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.login-header h1 {
  color: #333;
  margin-bottom: 10px;
  font-size: 28px;
}

.login-header p {
  color: #666;
  font-size: 14px;
}

.login-form {
  margin-bottom: 20px;
}

.login-tips {
  text-align: center;
  color: #999;
  font-size: 12px;
  line-height: 1.6;
}

.login-tips p {
  margin: 2px 0;
}
</style> 