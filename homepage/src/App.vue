<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth'
import Navbar from './components/Navbar.vue'
import Sidebar from './components/Sidebar.vue'
import Footer from './components/Footer.vue'

const router = useRouter()
const authStore = useAuthStore()
const isSidebarOpen = ref(false)

onMounted(async () => {
  // 检查用户认证状态
  await authStore.checkAuth()
})
</script>

<template>
  <div class="flex h-screen overflow-hidden">
    <!-- 侧边栏 -->
    <Sidebar
      :isOpen="isSidebarOpen"
      @close="isSidebarOpen = false"
    />

    <!-- 主容器 -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- 导航栏 -->
      <Navbar @menuClick="isSidebarOpen = !isSidebarOpen" />

      <!-- 主内容区 -->
      <div class="flex-1 overflow-y-auto">
        <RouterView />
      </div>

      <!-- 页脚 -->
      <Footer />
    </div>
  </div>
</template>

<style scoped></style>

