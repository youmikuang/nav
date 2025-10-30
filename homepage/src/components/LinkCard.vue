<script setup lang="ts">
import type { NavItem } from '../stores/nav'
import { ref } from 'vue'

interface Props {
  item: NavItem
}

defineProps<Props>()

const isHovered = ref(false)
const imageError = ref(false)

const handleImageError = () => {
  imageError.value = true
}
</script>

<template>
  <div
    class="group relative"
    @mouseenter="isHovered = true"
    @mouseleave="isHovered = false"
  >
    <!-- Hover Gradient Background -->
    <Transition name="fade">
      <div
        v-if="isHovered"
        class="absolute -inset-0.5 rounded-2xl bg-gradient-to-r from-blue-500/20 via-purple-500/20 to-pink-500/20 blur-sm"
      />
    </Transition>

    <a
      :href="item.link"
      target="_blank"
      rel="noopener noreferrer"
      class="relative z-10 block h-full overflow-hidden rounded-xl border border-gray-200 bg-white/90 p-5 shadow-sm transition-all duration-300 dark:border-gray-800 dark:bg-gray-900/80 hover:shadow-md dark:hover:shadow-lg"
      :class="{ 'scale-[1.02] shadow-lg dark:shadow-blue-500/10': isHovered }"
    >
      <!-- Icon and Title -->
      <div class="mb-3 flex items-center gap-3">
        <div class="relative h-10 w-10 flex-shrink-0">
          <img
            v-if="!imageError"
            :src="item.icon"
            :alt="item.title"
            class="h-full w-full rounded-lg object-cover"
            @error="handleImageError"
          />
          <div
            v-else
            class="h-full w-full rounded-lg bg-gray-200 dark:bg-gray-700"
          />
        </div>
        <h3 class="flex-1 truncate text-base font-semibold text-gray-900 dark:text-gray-100">
          {{ item.title }}
        </h3>
      </div>

      <!-- Description -->
      <p class="line-clamp-2 text-sm leading-relaxed text-gray-600 dark:text-gray-400">
        {{ item.desc }}
      </p>
    </a>
  </div>
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
</style>
