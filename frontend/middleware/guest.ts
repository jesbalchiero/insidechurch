import { useAuthStore } from '~/stores/auth'

export default defineNuxtRouteMiddleware((to) => {
  const authStore = useAuthStore()

  if (authStore.token) {
    return navigateTo('/dashboard')
  }
}) 