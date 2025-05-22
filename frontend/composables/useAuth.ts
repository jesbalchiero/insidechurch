import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'
import type { LoginRequest, RegisterRequest, User } from '~/types'

export const useAuth = () => {
  const router = useRouter()
  const toast = useToast()
  const user = ref<User | null>(null)
  const loading = ref(false)

  const login = async (credentials: LoginRequest) => {
    try {
      loading.value = true
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credentials)
      })

      if (!response.ok) {
        const error = await response.json()
        throw new Error(error.message || 'Erro ao fazer login')
      }

      const data = await response.json()
      user.value = data.user
      localStorage.setItem('token', data.token)
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
      const response = await fetch('/api/auth/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
      })

      if (!response.ok) {
        const error = await response.json()
        throw new Error(error.message || 'Erro ao cadastrar')
      }

      toast.success('Cadastro realizado com sucesso!')
      router.push('/login')
    } catch (error) {
      toast.error(error instanceof Error ? error.message : 'Erro ao cadastrar')
      throw error
    } finally {
      loading.value = false
    }
  }

  const logout = () => {
    user.value = null
    localStorage.removeItem('token')
    router.push('/login')
  }

  return {
    user,
    loading,
    login,
    register,
    logout
  }
} 