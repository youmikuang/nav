import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

interface User {
  id: number
  username: string
  email: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => user.value !== null)

  const apiClient = axios.create({
    baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api',
    withCredentials: true,
  })

  // 登录
  const login = async (username: string, password: string) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await apiClient.post('/auth/login', { username, password })
      user.value = response.data.user
      localStorage.setItem('token', response.data.token)
      apiClient.defaults.headers.common['Authorization'] = `Bearer ${response.data.token}`
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || '登录失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  // 登出
  const logout = async () => {
    isLoading.value = true
    try {
      await apiClient.post('/auth/logout')
      user.value = null
      localStorage.removeItem('token')
      delete apiClient.defaults.headers.common['Authorization']
    } catch (err: any) {
      error.value = err.response?.data?.message || '登出失败'
    } finally {
      isLoading.value = false
    }
  }

  // 检查认证状态
  const checkAuth = async () => {
    const token = localStorage.getItem('token')
    if (!token) {
      user.value = null
      return
    }

    try {
      apiClient.defaults.headers.common['Authorization'] = `Bearer ${token}`
      const response = await apiClient.get('/auth/me')
      user.value = response.data.user
    } catch (err) {
      user.value = null
      localStorage.removeItem('token')
      delete apiClient.defaults.headers.common['Authorization']
    }
  }

  // 注册
  const register = async (username: string, email: string, password: string) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await apiClient.post('/auth/register', { username, email, password })
      user.value = response.data.user
      localStorage.setItem('token', response.data.token)
      apiClient.defaults.headers.common['Authorization'] = `Bearer ${response.data.token}`
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || '注册失败'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    user,
    isLoading,
    error,
    isAuthenticated,
    login,
    logout,
    register,
    checkAuth,
  }
})
