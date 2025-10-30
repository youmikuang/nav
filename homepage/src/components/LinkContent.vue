<script setup lang="ts">
import { onMounted } from 'vue'
import { useNavStore } from '../stores/nav'
import LinkCard from './LinkCard.vue'

const navStore = useNavStore()

onMounted(() => {
  navStore.fetchCategories()
})
</script>

<template>
  <div class="min-h-screen w-full pb-32 pt-8">
    <div id="main" class="mx-auto w-full max-w-7xl px-6 md:px-8">
      <!-- Loading state -->
      <div
        v-if="navStore.isLoading"
        class="flex items-center justify-center py-12"
      >
        <div class="text-center">
          <div class="inline-block h-8 w-8 animate-spin rounded-full border-b-2 border-blue-600"></div>
          <p class="mt-4 text-gray-500 dark:text-gray-400">加载中...</p>
        </div>
      </div>

      <!-- Content -->
      <template v-else>
        <div
          v-for="(category, index) in navStore.categories"
          :id="`category-${index}`"
          :key="index"
          class="mb-16 scroll-mt-20"
        >
          <!-- Category title with gradient divider -->
          <div class="mb-6">
            <div class="flex items-center gap-4">
              <h2 class="text-3xl font-bold text-gray-900 dark:text-gray-50">
                {{ category.title }}
              </h2>
            </div>
          </div>

          <!-- Links grid - 3 columns on lg, 2 on md -->
          <div class="grid grid-cols-1 gap-5 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
            <LinkCard
              v-for="(item, itemIndex) in category.items"
              :key="itemIndex"
              :item="item"
            />
          </div>
        </div>
      </template>

      <!-- Error state -->
      <div
        v-if="navStore.error && !navStore.isLoading"
        class="rounded-lg border border-red-200 bg-red-50 p-4 text-red-600 dark:border-red-800 dark:bg-red-900/20 dark:text-red-400"
      >
        {{ navStore.error }}
      </div>
    </div>
  </div>
</template>

<style scoped></style>
