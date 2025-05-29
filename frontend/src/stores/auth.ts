import { defineStore } from 'pinia'
import type { AuthResponse } from '@/types/auth'

interface AuthState {
  token: string | null
  user: AuthResponse['user'] | null
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    token: localStorage.getItem('token'),
    user: JSON.parse(localStorage.getItem('user') || 'null')
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    currentUser: (state) => state.user
  },

  actions: {
    setAuth(auth: AuthResponse) {
      this.token = auth.token
      this.user = auth.user
      
      localStorage.setItem('token', auth.token)
      localStorage.setItem('user', JSON.stringify(auth.user))
    },

    clearAuth() {
      this.token = null
      this.user = null
      
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    }
  }
}) 