import { useAuthStore } from '~/stores/auth'

export default defineNuxtRouteMiddleware((to) => {
  const auth = useAuthStore()

  // Se a rota requer autenticação e o usuário não está autenticado
  if (to.meta.requiresAuth && !auth.token) {
    return navigateTo('/login')
  }

  // Se o usuário está autenticado e tenta acessar páginas de autenticação
  if (auth.token && (to.path === '/login' || to.path === '/register')) {
    return navigateTo('/')
  }
}) 