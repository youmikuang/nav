<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { Menu, LogOut, LogIn, Search, Github, Moon, Sun, Settings } from 'lucide-vue-next'

defineProps<{
  onMenuClick: () => void
}>()

const router = useRouter()
const authStore = useAuthStore()
const userMenuOpen = ref(false)
const isDarkMode = ref(false)

// Check system preference on mount
const initializeDarkMode = () => {
  const isDark = document.documentElement.classList.contains('dark')
  isDarkMode.value = isDark
}

onMounted(() => {
  initializeDarkMode()
})

const toggleDarkMode = () => {
  isDarkMode.value = !isDarkMode.value
  if (isDarkMode.value) {
    document.documentElement.classList.add('dark')
    localStorage.setItem('theme', 'dark')
  } else {
    document.documentElement.classList.remove('dark')
    localStorage.setItem('theme', 'light')
  }
}

const handleLogout = async () => {
  await authStore.logout()
  router.push('/')
  userMenuOpen.value = false
}

const handleLogin = () => {
  router.push('/login')
  userMenuOpen.value = false
}

const handleRegister = () => {
  router.push('/register')
  userMenuOpen.value = false
}

const handleSettings = () => {
  // TODO: Navigate to settings page
  console.log('Settings')
  userMenuOpen.value = false
}
</script>

<template>
  <header class="sticky top-0 z-40 w-full border-b border-gray-200 bg-white dark:border-gray-800 dark:bg-slate-950">
    <div class="container mx-auto flex h-16 items-center justify-between gap-4 px-4">

      <!-- Center: Search -->
      <div class="hidden flex-1 mx-4 sm:block lg:mx-8">
        <div class="relative">
          <Search class="absolute left-3 top-2.5 h-5 w-5 text-gray-400" />
          <input
            type="text"
            placeholder="搜索网站..."
            class="w-full max-w-sm rounded-lg border border-gray-300 bg-white py-2 pl-10 pr-4 text-sm transition-all focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 dark:border-gray-700 dark:bg-slate-900 dark:text-gray-100 dark:focus:border-blue-400 dark:focus:ring-blue-400"
          />
        </div>
      </div>

      <!-- Right: Actions -->
      <nav class="flex items-center gap-2 sm:gap-4">
        <!-- GitHub icon -->
        <a
          href="https://github.com"
          target="_blank"
          rel="noopener noreferrer"
          class="p-2 hover:bg-gray-100 rounded-lg dark:hover:bg-slate-800 transition-colors"
          title="GitHub"
        >
          <Github class="h-5 w-5" />
        </a>

        <!-- Dark mode toggle -->
        <button
          class="p-2 hover:bg-gray-100 rounded-lg dark:hover:bg-slate-800 transition-colors"
          @click="toggleDarkMode"
          title="Toggle dark mode"
        >
          <Moon v-if="isDarkMode" class="h-5 w-5" />
          <Sun v-else class="h-5 w-5" />
        </button>

        <!-- Settings icon -->
        <button
          class="p-2 hover:bg-gray-100 rounded-lg dark:hover:bg-slate-800 transition-colors sm:hidden"
          @click="handleSettings"
          title="Settings"
        >
          <Settings class="h-5 w-5" />
        </button>

        <!-- Mobile menu button -->
        <button
          class="sm:hidden p-2 hover:bg-gray-100 rounded-lg dark:hover:bg-slate-800"
          @click="$emit('menuClick')"
        >
          <Menu class="h-5 w-5" />
        </button>

        <!-- Auth buttons / User menu -->
        <div class="relative">
          <template v-if="authStore.isAuthenticated">
            <button
              class="p-2 hover:bg-gray-100 rounded-lg dark:hover:bg-slate-800"
              @click="userMenuOpen = !userMenuOpen"
            >
              <div class="w-8 h-8 bg-primary-600 rounded-full flex items-center justify-center text-white text-sm font-bold">
                {{ authStore.user?.username?.charAt(0).toUpperCase() }}
              </div>
            </button>

            <!-- User dropdown menu -->
            <div
              v-if="userMenuOpen"
              class="absolute right-0 mt-2 w-48 bg-white dark:bg-slate-800 rounded-lg shadow-lg border border-gray-200 dark:border-gray-700"
            >
              <div class="px-4 py-3 border-b border-gray-200 dark:border-gray-700">
                <p class="font-semibold text-sm">{{ authStore.user?.username }}</p>
                <p class="text-xs text-gray-500">{{ authStore.user?.email }}</p>
              </div>
              <button
                class="w-full text-left px-4 py-2 hover:bg-gray-100 dark:hover:bg-slate-700 flex items-center gap-2 text-red-600"
                @click="handleLogout"
              >
                <LogOut class="h-4 w-4" />
                登出
              </button>
            </div>
          </template>

          <template v-else>
            <div class="hidden sm:flex items-center gap-2">
              <button
                class="px-4 py-2 text-sm font-medium text-primary-600 hover:bg-primary-50 rounded-lg dark:hover:bg-slate-800"
                @click="handleLogin"
              >
                <LogIn class="inline h-4 w-4 mr-1" />
                登录
              </button>
              <button
                class="px-4 py-2 text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 rounded-lg"
                @click="handleRegister"
              >
                注册
              </button>
            </div>
          </template>
        </div>
      </nav>
    </div>
  </header>
</template>

<style scoped></style>
