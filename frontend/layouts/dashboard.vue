<template>
  <div class="min-h-screen bg-base-200">
    <!-- Header -->
    <header class="bg-base-100 shadow-sm">
      <div class="container mx-auto px-4">
        <div class="flex items-center justify-between h-16">
          <!-- Logo e Nome -->
          <div class="flex items-center space-x-4">
            <button
              class="md:hidden btn btn-ghost btn-circle"
              @click="isMenuOpen = !isMenuOpen"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  v-if="isMenuOpen"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"
                />
                <path
                  v-else
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 6h16M4 12h16M4 18h16"
                />
              </svg>
            </button>
            <NuxtLink to="/dashboard" class="text-xl font-bold">
              InsideChurch
            </NuxtLink>
          </div>

          <!-- Menu Desktop -->
          <div class="hidden md:flex items-center space-x-4">
            <span class="text-base-content/80">{{ user?.value?.name }}</span>
            <button
              class="btn btn-ghost btn-sm"
              @click="handleLogout"
            >
              Sair
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Menu Mobile -->
    <div
      v-if="isMenuOpen"
      class="md:hidden fixed inset-0 z-50 bg-base-100"
    >
      <div class="flex flex-col h-full">
        <div class="flex items-center justify-between p-4 border-b">
          <span class="text-xl font-bold">Menu</span>
          <button
            class="btn btn-ghost btn-circle"
            @click="isMenuOpen = false"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </div>

        <div class="flex-1 p-4">
          <div class="space-y-4">
            <div class="text-base-content/80">
              {{ user?.value?.name }}
            </div>
            <button
              class="btn btn-ghost w-full justify-start"
              @click="handleLogout"
            >
              Sair
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Conteúdo Principal -->
    <main class="container mx-auto px-4 py-8">
      <slot />
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
</script> 