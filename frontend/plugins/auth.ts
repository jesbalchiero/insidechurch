import { useAuthStore } from '~/stores/auth'

export default defineNuxtPlugin(() => {
  const authStore = useAuthStore()

  // Inicializa o estado de autenticação
  authStore.initialize()

  // Adiciona interceptor para atualizar token
  const { $fetch } = useNuxtApp()
  $fetch.interceptors.response.use(
    (response) => response,
    async (error) => {
      const originalRequest = error.config

      if (error.response?.status === 401 && !originalRequest._retry) {
        originalRequest._retry = true

        try {
          await authStore.refreshAccessToken()
          originalRequest.headers['Authorization'] = `Bearer ${authStore.token}`
          return $fetch(originalRequest)
        } catch (refreshError) {
          authStore.logout()
          return Promise.reject(refreshError)
        }
      }

      return Promise.reject(error)
    }
  )
}) 