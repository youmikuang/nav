import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

export interface NavItem {
  icon: string
  title: string
  desc: string
  link: string
}

export interface NavCategory {
  id?: number
  title: string
  items: NavItem[]
}

export const useNavStore = defineStore('nav', () => {
  const categories = ref<NavCategory[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const apiClient = axios.create({
    baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api',
    withCredentials: true,
  })

  // 从后端获取导航数据
  const fetchCategories = async () => {
    isLoading.value = true
    error.value = null
    try {
      const response = await apiClient.get('/nav/categories')
      categories.value = response.data.categories || response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || '获取导航数据失败'
      console.error('Failed to fetch categories:', err)
      // 使用默认数据
      loadDefaultCategories()
    } finally {
      isLoading.value = false
    }
  }

  // 加载默认分类
  const loadDefaultCategories = () => {
    categories.value = [
      {
        title: '个人网站',
        items: [
          {
            icon: 'https://todo.gaoheng.top/assets/favicon.png',
            title: 'TODO List',
            desc: '一个简单好用的待办事项应用',
            link: 'https://todo.gaoheng.top/',
          },
        ],
      },
      {
        title: '常用工具',
        items: [
          {
            icon: 'https://caniuse.com/img/favicon-128.png',
            title: 'Can I use',
            desc: '前端 API 兼容性查询',
            link: 'https://caniuse.com',
          },
        ],
      },
    ]
  }

  return {
    categories,
    isLoading,
    error,
    fetchCategories,
    loadDefaultCategories,
  }
})
