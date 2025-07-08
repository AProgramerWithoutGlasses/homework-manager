import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

export const useAuthStore = defineStore('auth', () => {
  const token = ref(localStorage.getItem('token') || '')
  let teacherRaw = localStorage.getItem('teacher')
  if (teacherRaw === null || teacherRaw === undefined || teacherRaw === 'undefined') {
    teacherRaw = null
  }
  const teacher = ref(teacherRaw ? JSON.parse(teacherRaw) : null)

  const isAuthenticated = computed(() => !!token.value)

  const login = async (telephone, password) => {
    try {
      const response = await axios.post('/backend/login', {
        telephone,
        password
      })
      
      const { token: newToken, teacher: teacherData } = response.data
      
      // 更新状态
      token.value = newToken
      teacher.value = teacherData
      
      // 保存到localStorage
      localStorage.setItem('token', newToken)
      localStorage.setItem('teacher', JSON.stringify(teacherData))
      
      // 设置axios默认headers
      axios.defaults.headers.common['Authorization'] = `Bearer ${newToken}`
      
      return { success: true }
    } catch (error) {
      return { 
        success: false, 
        error: error.response?.data?.error || '登录失败' 
      }
    }
  }

  const logout = () => {
    token.value = ''
    teacher.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('teacher')
    delete axios.defaults.headers.common['Authorization']
  }

  const initAuth = () => {
    // 从localStorage重新读取状态
    const storedToken = localStorage.getItem('token')
    const storedTeacher = localStorage.getItem('teacher')
    
    if (storedToken) {
      token.value = storedToken
      axios.defaults.headers.common['Authorization'] = `Bearer ${storedToken}`
    }
    
    if (storedTeacher && storedTeacher !== 'undefined') {
      try {
        teacher.value = JSON.parse(storedTeacher)
      } catch (e) {
        teacher.value = null
      }
    }
  }

  return {
    token,
    teacher,
    isAuthenticated,
    login,
    logout,
    initAuth
  }
}) 