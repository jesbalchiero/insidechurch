import type { User, LoginRequest, RegisterRequest } from '~/types'
import { useApi } from './useApi'
import { useToast } from './useToast'

export const useAuth = () => {
  const api = useApi()
  const toast = useToast()
  const user = useState<User | null>('user', () => null)

  const login = async (email: string, password: string) => {
    try {
      const response = await api.fetch<{ token: string }>('/api/login', {
        method: 'POST',
        body: { email, password } as LoginRequest,
      })

      localStorage.setItem('token', response.token)
      await getUser()
      toast.success('Login realizado com sucesso!')
      return response
    } catch (error: any) {
      toast.error(error.data?.message || 'Erro ao fazer login')
      throw error
    }
  }

  const register = async (email: string, password: string, name: string) => {
    try {
      const response = await api.fetch<{ token: string }>('/api/register', {
        method: 'POST',
        body: { email, password, name } as RegisterRequest,
      })

      localStorage.setItem('token', response.token)
      await getUser()
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
      user.value = response
      return response
    } catch (error: any) {
      if (error.statusCode === 401) {
        user.value = null
      }
      throw error
    }
  }

  const logout = () => {
    localStorage.removeItem('token')
    user.value = null
    toast.info('Logout realizado com sucesso!')
  }

  return {
    user,
    login,
    register,
    getUser,
    logout,
  }
} 