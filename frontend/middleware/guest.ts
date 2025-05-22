import { useAuth } from '~/composables/useAuth'

export default defineNuxtRouteMiddleware((to) => {
  const auth = useAuth()
  const token = localStorage.getItem('token')

  if (token) {
    // Se tiver token mas não tiver usuário, tenta carregar
    if (!auth.user.value) {
      try {
        auth.getUser()
        return navigateTo('/dashboard')
      } catch (error) {
        // Se der erro, limpa autenticação e permite acesso
        auth.logout()
      }
    } else {
      // Se já tiver usuário carregado, redireciona
      return navigateTo('/dashboard')
    }
  }
}) 