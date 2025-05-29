import { defineStore } from 'pinia'
import { ref } from 'vue'

interface User {
  id: string
  email: string
}

interface LoginResponse {
  token: string
  refreshToken: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)
  const refreshToken = ref<string | null>(null)

  const register = async (email: string, password: string) => {
    try {
      const response = await fetch('http://localhost:8081/auth/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })

      if (!response.ok) {
        const error = await response.text()
        throw new Error(error)
      }

      const data: LoginResponse = await response.json()
      token.value = data.token
      refreshToken.value = data.refreshToken

      // Decodificar o token para obter informações do usuário
      const payload = JSON.parse(atob(data.token.split('.')[1]))
      user.value = {
        id: payload.user_id,
        email: email,
      }

      // Salvar no localStorage
      localStorage.setItem('token', data.token)
      localStorage.setItem('refreshToken', data.refreshToken)
    } catch (error) {
      console.error('Erro ao registrar:', error)
      throw error
    }
  }

  const login = async (email: string, password: string) => {
    try {
      const response = await fetch('http://localhost:8081/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })

      if (!response.ok) {
        const error = await response.text()
        throw new Error(error)
      }

      const data: LoginResponse = await response.json()
      token.value = data.token
      refreshToken.value = data.refreshToken

      // Decodificar o token para obter informações do usuário
      const payload = JSON.parse(atob(data.token.split('.')[1]))
      user.value = {
        id: payload.user_id,
        email: email,
      }

      // Salvar no localStorage
      localStorage.setItem('token', data.token)
      localStorage.setItem('refreshToken', data.refreshToken)
    } catch (error) {
      console.error('Erro ao fazer login:', error)
      throw error
    }
  }

  const logout = () => {
    user.value = null
    token.value = null
    refreshToken.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('refreshToken')
  }

  const refreshAccessToken = async () => {
    if (!refreshToken.value) {
      throw new Error('Refresh token não disponível')
    }

    try {
      const response = await fetch('http://localhost:8081/auth/refresh', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ refresh_token: refreshToken.value }),
      })

      if (!response.ok) {
        throw new Error('Erro ao atualizar token')
      }

      const data: LoginResponse = await response.json()
      token.value = data.token
      refreshToken.value = data.refreshToken

      // Atualizar no localStorage
      localStorage.setItem('token', data.token)
      localStorage.setItem('refreshToken', data.refreshToken)
    } catch (error) {
      console.error('Erro ao atualizar token:', error)
      logout()
      throw error
    }
  }

  const initialize = () => {
    const storedToken = localStorage.getItem('token')
    const storedRefreshToken = localStorage.getItem('refreshToken')

    if (storedToken && storedRefreshToken) {
      token.value = storedToken
      refreshToken.value = storedRefreshToken

      // Decodificar o token para obter informações do usuário
      try {
        const payload = JSON.parse(atob(storedToken.split('.')[1]))
        user.value = {
          id: payload.user_id,
          email: payload.email,
        }
      } catch (error) {
        console.error('Erro ao decodificar token:', error)
        logout()
      }
    }
  }

  return {
    user,
    token,
    refreshToken,
    register,
    login,
    logout,
    refreshAccessToken,
    initialize,
  }
}) 