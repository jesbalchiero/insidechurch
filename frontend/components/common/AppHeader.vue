<template>
  <header class="bg-white shadow-sm">
    <nav class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <div class="flex h-16 justify-between items-center">
        <!-- Logo -->
        <div class="flex-shrink-0">
          <NuxtLink to="/" class="text-2xl font-bold text-primary-600">
            InsideChurch
          </NuxtLink>
        </div>

        <!-- Desktop Navigation -->
        <div class="hidden md:flex md:items-center md:space-x-4">
          <template v-if="!auth.user">
            <NuxtLink
              to="/login"
              class="text-gray-600 hover:text-primary-600 px-3 py-2 rounded-md text-sm font-medium"
            >
              Entrar
            </NuxtLink>
            <NuxtLink
              to="/register"
              class="bg-primary-600 text-white hover:bg-primary-700 px-4 py-2 rounded-md text-sm font-medium"
            >
              Cadastrar
            </NuxtLink>
          </template>

          <template v-else>
            <Menu as="div" class="relative">
              <MenuButton
                class="flex items-center space-x-2 text-gray-700 hover:text-primary-600 focus:outline-none"
              >
                <div class="h-8 w-8 rounded-full bg-primary-100 flex items-center justify-center">
                  <span class="text-sm font-medium text-primary-600">
                    {{ userInitials }}
                  </span>
                </div>
              </MenuButton>

              <transition
                enter-active-class="transition ease-out duration-100"
                enter-from-class="transform opacity-0 scale-95"
                enter-to-class="transform opacity-100 scale-100"
                leave-active-class="transition ease-in duration-75"
                leave-from-class="transform opacity-100 scale-100"
                leave-to-class="transform opacity-0 scale-95"
              >
                <MenuItems
                  class="absolute right-0 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
                >
                  <div class="py-1">
                    <div class="px-4 py-2 text-sm text-gray-700 border-b">
                      <div class="font-medium">{{ auth.user.value?.name }}</div>
                      <div class="text-gray-500">{{ auth.user.value?.email }}</div>
                    </div>
                    <MenuItem v-slot="{ active }">
                      <button
                        @click="auth.logout"
                        :class="[
                          active ? 'bg-gray-100' : '',
                          'w-full text-left px-4 py-2 text-sm text-gray-700'
                        ]"
                      >
                        Sair
                      </button>
                    </MenuItem>
                  </div>
                </MenuItems>
              </transition>
            </Menu>
          </template>
        </div>

        <!-- Mobile menu button -->
        <div class="flex md:hidden">
          <button
            @click="isOpen = !isOpen"
            class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-primary-500"
          >
            <span class="sr-only">Abrir menu</span>
            <Icon
              :name="isOpen ? 'heroicons:x-mark' : 'heroicons:bars-3'"
              class="h-6 w-6"
            />
          </button>
        </div>
      </div>
    </nav>

    <!-- Mobile drawer -->
    <Transition
      enter-active-class="transition-opacity ease-linear duration-300"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition-opacity ease-linear duration-300"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-if="isOpen"
        class="fixed inset-0 bg-gray-600 bg-opacity-75 z-40"
        @click="isOpen = false"
      ></div>
    </Transition>

    <Transition
      enter-active-class="transition ease-in-out duration-300 transform"
      enter-from-class="-translate-x-full"
      enter-to-class="translate-x-0"
      leave-active-class="transition ease-in-out duration-300 transform"
      leave-from-class="translate-x-0"
      leave-to-class="-translate-x-full"
    >
      <div
        v-if="isOpen"
        class="fixed inset-y-0 left-0 w-64 bg-white shadow-lg z-50"
      >
        <div class="h-full flex flex-col">
          <div class="px-4 py-6 border-b">
            <h2 class="text-lg font-medium text-gray-900">Menu</h2>
          </div>

          <div class="flex-1 px-4 py-6 space-y-4">
            <template v-if="!auth.user">
              <NuxtLink
                to="/login"
                class="block text-gray-600 hover:text-primary-600"
                @click="isOpen = false"
              >
                Entrar
              </NuxtLink>
              <NuxtLink
                to="/register"
                class="block text-gray-600 hover:text-primary-600"
                @click="isOpen = false"
              >
                Cadastrar
              </NuxtLink>
            </template>

            <template v-else>
              <div class="py-2">
                <div class="font-medium text-gray-900">{{ auth.user.value?.name }}</div>
                <div class="text-sm text-gray-500">{{ auth.user.value?.email }}</div>
              </div>
              <button
                @click="handleLogout"
                class="w-full text-left text-gray-600 hover:text-primary-600"
              >
                Sair
              </button>
            </template>
          </div>
        </div>
      </div>
    </Transition>
  </header>
</template>

<script setup lang="ts">
import { Menu, MenuButton, MenuItems, MenuItem } from '@headlessui/vue'
import { useAuth } from '~/composables/useAuth'

const auth = useAuth()
const isOpen = ref(false)

const userInitials = computed(() => {
  if (!auth.user.value?.name) return ''
  return auth.user.value.name
    .split(' ')
    .map(n => n[0])
    .join('')
    .toUpperCase()
})

const handleLogout = () => {
  auth.logout()
  isOpen.value = false
}
</script> 