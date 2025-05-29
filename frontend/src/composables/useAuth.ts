import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'
import { useApi } from './useApi'
import { useAuthStore } from '@/stores/auth'
import type { LoginRequest, RegisterRequest, AuthResponse, AuthError } from '@/types/auth'

export const useAuth = () => {
  const { fetch } = useApi()
  const router = useRouter()
  const toast = useToast()
  const authStore = useAuthStore()

  const handleAuthError = (error: any) => {
    const authError = error as AuthError
    
    switch (authError.code) {
      case 'INVALID_CREDENTIALS':
        toast.error('Email ou senha inválidos')
        break
      case 'EMAIL_EXISTS':
        toast.error('Este email já está em uso')
        break
      case 'UNAUTHORIZED':
        toast.error('Sessão expirada. Por favor, faça login novamente')
        authStore.clearAuth()
        router.push('/login')
        break
      default:
        toast.error(authError.message || 'Ocorreu um erro. Tente novamente.')
    }
  }

  const login = async (credentials: LoginRequest) => {
    try {
      const response = await fetch<AuthResponse>('/auth/login', {
        method: 'POST',
        body: credentials
      })
      
      authStore.setAuth(response)
      toast.success('Login realizado com sucesso!')
      await router.push('/dashboard')
    } catch (error) {
      handleAuthError(error)
      throw error
    }
  }

  const register = async (data: RegisterRequest) => {
    try {
      const response = await fetch<AuthResponse>('/auth/register', {
        method: 'POST',
        body: data
      })
      
      authStore.setAuth(response)
      toast.success('Cadastro realizado com sucesso!')
      await router.push('/dashboard')
    } catch (error) {
      handleAuthError(error)
      throw error
    }
  }

  const logout = async () => {
    try {
      authStore.clearAuth()
      toast.success('Logout realizado com sucesso!')
      await router.push('/login')
    } catch (error) {
      toast.error('Erro ao realizar logout')
      throw error
    }
  }

  const checkAuth = async () => {
    try {
      const response = await fetch<AuthResponse>('/users/me')
      authStore.setAuth(response)
      return true
    } catch (error) {
      authStore.clearAuth()
      return false
    }
  }

  return {
    login,
    register,
    logout,
    checkAuth
  }
} 