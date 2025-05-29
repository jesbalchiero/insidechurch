import { useAuth } from '~/composables/useAuth'

export default defineNuxtRouteMiddleware((to) => {
  const auth = useAuth()
  let token = null
  
  if (process.client) {
    token = localStorage.getItem('token')
  }

  if (!token) {
    return navigateTo('/login')
  }

  // Se não tiver usuário carregado, tenta carregar
  if (!auth.user.value) {
    try {
      auth.getUser()
    } catch (error) {
      // Se der erro, limpa autenticação e redireciona
      auth.logout()
      return navigateTo('/login')
    }
  }
}) 