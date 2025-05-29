import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from './useToast'
import { useApi } from './useApi'
import type { LoginRequest, RegisterRequest, User } from '~/types'
import { useAuthStore } from '~/stores/auth'

export const useAuth = () => {
  const router = useRouter()
  const toast = useToast()
  const api = useApi()
  const authStore = useAuthStore()
  const user = ref<User | null>(null)
  const loading = ref(false)

  const isAuthenticated = computed(() => !!authStore.token)

  const login = async (email: string, password: string) => {
    await authStore.login(email, password)
  }

  const register = async (email: string, password: string) => {
    await authStore.register(email, password)
  }

  const logout = () => {
    authStore.logout()
  }

  const refreshToken = async () => {
    await authStore.refreshAccessToken()
  }

  const getUser = async () => {
    try {
      loading.value = true
      let token = null
      if (process.client) {
        token = localStorage.getItem('token')
      }
      
      if (!token) {
        throw new Error('Token não encontrado')
      }

      const data = await api.fetch<User>('/user')
      user.value = data
      return data
    } catch (error: any) {
      if (error.statusCode === 401) {
        if (process.client) {
          localStorage.removeItem('token')
        }
        router.push('/login')
      }
      toast.error(error instanceof Error ? error.message : 'Erro ao carregar usuário')
      throw error
    } finally {
      loading.value = false
    }
  }

  return {
    user,
    loading,
    isAuthenticated,
    login,
    register,
    getUser,
    logout,
    refreshToken,
  }
} 