<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { Mail, Lock, User, ArrowLeft } from 'lucide-vue-next'

const router = useRouter()
const authStore = useAuthStore()

const username = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')

const handleRegister = async () => {
  if (!username.value || !email.value || !password.value || !confirmPassword.value) {
    alert('请填写所有字段')
    return
  }

  if (password.value !== confirmPassword.value) {
    alert('两次输入的密码不一致')
    return
  }

  if (password.value.length < 6) {
    alert('密码长度至少6位')
    return
  }

  try {
    await authStore.register(username.value, email.value, password.value)
    router.push('/')
  } catch (error) {
    // Error is handled in store
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-primary-50 to-blue-50 dark:from-slate-950 dark:to-slate-900 px-4">
    <div class="w-full max-w-md">
      <!-- Back button -->
      <button
        @click="$router.back()"
        class="flex items-center gap-2 mb-6 text-primary-600 hover:text-primary-700 dark:text-primary-400"
      >
        <ArrowLeft class="h-4 w-4" />
        返回
      </button>

      <!-- Card -->
      <div class="bg-white dark:bg-slate-900 rounded-lg shadow-xl p-8">
        <!-- Title -->
        <h1 class="text-3xl font-bold text-center mb-2 text-gray-900 dark:text-white">
          注册
        </h1>
        <p class="text-center text-gray-600 dark:text-gray-400 mb-8">
          创建一个新账户，开始探索
        </p>

        <!-- Form -->
        <form @submit.prevent="handleRegister" class="space-y-6">
          <!-- Username -->
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              用户名
            </label>
            <div class="relative">
              <User class="absolute left-3 top-3 h-5 w-5 text-gray-400" />
              <input
                v-model="username"
                type="text"
                placeholder="输入用户名"
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-600 dark:bg-slate-800 dark:border-gray-700 dark:text-white"
              />
            </div>
          </div>

          <!-- Email -->
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              邮箱
            </label>
            <div class="relative">
              <Mail class="absolute left-3 top-3 h-5 w-5 text-gray-400" />
              <input
                v-model="email"
                type="email"
                placeholder="输入邮箱"
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-600 dark:bg-slate-800 dark:border-gray-700 dark:text-white"
              />
            </div>
          </div>

          <!-- Password -->
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              密码
            </label>
            <div class="relative">
              <Lock class="absolute left-3 top-3 h-5 w-5 text-gray-400" />
              <input
                v-model="password"
                type="password"
                placeholder="输入密码（至少6位）"
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-600 dark:bg-slate-800 dark:border-gray-700 dark:text-white"
              />
            </div>
          </div>

          <!-- Confirm Password -->
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              确认密码
            </label>
            <div class="relative">
              <Lock class="absolute left-3 top-3 h-5 w-5 text-gray-400" />
              <input
                v-model="confirmPassword"
                type="password"
                placeholder="确认密码"
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-600 dark:bg-slate-800 dark:border-gray-700 dark:text-white"
              />
            </div>
          </div>

          <!-- Error message -->
          <div
            v-if="authStore.error"
            class="rounded-lg border border-red-200 bg-red-50 p-3 text-red-600 text-sm dark:border-red-800 dark:bg-red-900/20 dark:text-red-400"
          >
            {{ authStore.error }}
          </div>

          <!-- Submit button -->
          <button
            type="submit"
            :disabled="authStore.isLoading"
            class="w-full bg-primary-600 hover:bg-primary-700 disabled:bg-gray-400 text-white font-medium py-2 rounded-lg transition-colors"
          >
            {{ authStore.isLoading ? '注册中...' : '注册' }}
          </button>
        </form>

        <!-- Divider -->
        <div class="my-6 flex items-center gap-4">
          <div class="flex-1 h-px bg-gray-300 dark:bg-gray-700"></div>
          <span class="text-sm text-gray-500 dark:text-gray-400">或</span>
          <div class="flex-1 h-px bg-gray-300 dark:bg-gray-700"></div>
        </div>

        <!-- Login link -->
        <p class="text-center text-gray-600 dark:text-gray-400">
          已有账户？
          <RouterLink
            to="/login"
            class="text-primary-600 hover:text-primary-700 font-medium"
          >
            立即登录
          </RouterLink>
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
