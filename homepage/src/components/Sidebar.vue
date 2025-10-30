<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useNavStore } from '../stores/nav'
import { X } from 'lucide-vue-next'

interface Props {
  isOpen: boolean
}

const props = defineProps<Props>()

const emit = defineEmits<{
  close: []
}>()

const navStore = useNavStore()
const activeTabId = ref(0)

const scroll = (index: number) => {
  const ele = document.getElementById(`category-${index}`)
  if (ele) {
    const elePosition = ele.getBoundingClientRect().top
    const offsetPosition = elePosition + window.scrollY - 80
    window.scrollTo({
      top: offsetPosition,
      behavior: 'smooth',
    })
  }
  emit('close')
}

const scrollUpdate = () => {
  const main = document.getElementById('main')
  if (main) {
    const children = Array.from(main.children)
    for (const child of children) {
      const top = child.getBoundingClientRect().top
      if (top < 100) {
        const id = child.id.replace('category-', '')
        activeTabId.value = Number(id)
      }
    }
  }
}

const hoveredIndex = ref<number | null>(null)

onMounted(() => {
  window.addEventListener('scroll', scrollUpdate)
})

onUnmounted(() => {
  window.removeEventListener('scroll', scrollUpdate)
})
</script>

<template>
  <!-- Mobile overlay -->
  <Transition name="fade">
    <div
      v-if="isOpen"
      class="fixed inset-0 z-30 bg-black/50 sm:hidden"
      @click="$emit('close')"
    />
  </Transition>

  <!-- Sidebar -->
  <nav
    :class="[
      'fixed z-40 h-screen w-64 border-r border-gray-200 bg-white/80 pt-16 backdrop-blur-sm transition-transform duration-300 ease-in-out dark:border-gray-800 dark:bg-gray-950/80 sm:static sm:translate-x-0 sm:pt-0',
      isOpen ? 'translate-x-0' : '-translate-x-full',
    ]"
  >
    <!-- Close button on mobile -->
    <button
      class="absolute right-4 top-4 rounded-lg p-2 hover:bg-gray-100 dark:hover:bg-gray-800 sm:hidden"
      @click="$emit('close')"
    >
      <X class="h-5 w-5" />
    </button>

    <!-- Sidebar header -->
    <div class="hidden border-b border-gray-200 px-6 py-6 dark:border-gray-800 sm:block">
      <h2 class="text-xl font-bold text-gray-900 dark:text-gray-50">网址导航</h2>
    </div>

    <!-- Category list -->
    <div class="overflow-y-auto px-4 py-6 sm:px-6">
      <div class="space-y-2">
        <div
          v-for="(category, index) in navStore.categories"
          :key="index"
          class="relative"
        >
          <!-- Active/Hover background animation -->
          <Transition name="fade-scale">
            <div
              v-if="activeTabId === index || hoveredIndex === index"
              class="absolute inset-0 rounded-lg bg-blue-50 dark:bg-blue-900/20"
            />
          </Transition>

          <button
            @click="scroll(index)"
            @mouseenter="hoveredIndex = index"
            @mouseleave="hoveredIndex = null"
            :class="[
              'relative z-10 w-full cursor-pointer rounded-lg px-4 py-2.5 text-left text-sm transition-all duration-200',
              activeTabId === index
                ? 'font-semibold text-blue-600 dark:text-blue-400'
                : 'text-gray-700 hover:text-gray-900 dark:text-gray-300 dark:hover:text-gray-100',
            ]"
          >
            <span class="truncate">{{ category.title }}</span>
          </button>
        </div>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.fade-scale-enter-active,
.fade-scale-leave-active {
  transition: opacity 0.2s ease;
}

.fade-scale-enter-from,
.fade-scale-leave-to {
  opacity: 0;
}
</style>
