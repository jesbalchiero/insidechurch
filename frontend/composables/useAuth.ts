import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from './useToast'
import { useApi } from './useApi'
import type { LoginRequest, RegisterRequest, User } from '~/types'

export const useAuth = () => {
  const router = useRouter()
  const toast = useToast()
  const api = useApi()
  const user = ref<User | null>(null)
  const loading = ref(false)

  const login = async (credentials: LoginRequest) => {
    try {
      loading.value = true
      const data = await api.fetch<{ user: User; token: string }>('/auth/login', {
        method: 'POST',
        body: credentials
      })

      user.value = data.user
      if (process.client) {
        localStorage.setItem('token', data.token)
      }
      toast.success('Login realizado com sucesso!')
      router.push('/dashboard')
    } catch (error) {
      toast.error(error instanceof Error ? error.message : 'Erro ao fazer login')
      throw error
    } finally {
      loading.value = false
    }
  }

  const register = async (data: RegisterRequest) => {
    try {
      loading.value = true
      await api.fetch('/auth/register', {
        method: 'POST',
        body: data
      })

      toast.success('Cadastro realizado com sucesso!')
      router.push('/login')
    } catch (error) {
      toast.error(error instanceof Error ? error.message : 'Erro ao cadastrar')
      throw error
    } finally {
      loading.value = false
    }
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

  const logout = () => {
    user.value = null
    if (process.client) {
      localStorage.removeItem('token')
    }
    router.push('/login')
  }

  return {
    user,
    loading,
    login,
    register,
    getUser,
    logout
  }
} 