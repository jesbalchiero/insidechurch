import { useAuth } from '~/composables/useAuth'

export default defineNuxtPlugin(async () => {
  const auth = useAuth()
  let token = null
  
  if (process.client) {
    token = localStorage.getItem('token')
  }

  if (token) {
    try {
      await auth.getUser()
    } catch (error) {
      // Se der erro 401 ou qualquer outro, limpa autenticação
      auth.logout()
    }
  }
}) 