import type { User, LoginRequest, RegisterRequest } from '~/types'
import { useApi } from './useApi'
import { useToast } from './useToast'
import { useAuthStore } from '~/stores/auth'

export const useAuth = () => {
  const api = useApi()
  const toast = useToast()
  const auth = useAuthStore()

  const login = async (email: string, password: string) => {
    try {
      const response = await api.fetch<{ token: string; user: User }>('/api/login', {
        method: 'POST',
        body: { email, password } as LoginRequest,
      })

      auth.setAuth(response)
      toast.success('Login realizado com sucesso!')
      return response
    } catch (error: any) {
      toast.error(error.data?.message || 'Erro ao fazer login')
      throw error
    }
  }

  const register = async (email: string, password: string, name: string) => {
    try {
      const response = await api.fetch<{ token: string; user: User }>('/api/register', {
        method: 'POST',
        body: { email, password, name } as RegisterRequest,
      })

      auth.setAuth(response)
      toast.success('Cadastro realizado com sucesso!')
      return response
    } catch (error: any) {
      toast.error(error.data?.message || 'Erro ao fazer cadastro')
      throw error
    }
  }

  const getUser = async () => {
    try {
      const response = await api.fetch<User>('/api/user')
      auth.setUser(response)
      return response
    } catch (error: any) {
      if (error.statusCode === 401) {
        auth.clearAuth()
      }
      throw error
    }
  }

  const logout = () => {
    auth.clearAuth()
    toast.info('Logout realizado com sucesso!')
  }

  return {
    user: auth.user,
    isAuthenticated: auth.isAuthenticated,
    login,
    register,
    getUser,
    logout,
  }
} 