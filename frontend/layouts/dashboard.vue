<template>
  <div class="min-h-screen bg-gray-50">
    <AppHeader />
    <main class="py-10">
      <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <slot />
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '~/composables/useAuth'

const auth = useAuth()
const user = computed(() => auth.user)
const isMenuOpen = ref(false)

const handleLogout = () => {
  auth.logout()
  isMenuOpen.value = false
}

definePageMeta({
  middleware: ['auth']
})
</script> 