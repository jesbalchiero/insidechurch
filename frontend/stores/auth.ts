import { defineStore } from 'pinia'
import type { User, AuthResponse } from '~/types'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value)

  function setAuth(response: AuthResponse) {
    token.value = response.token
    user.value = {
      ...response.user,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
    }
    
    // Salvar token no cookie
    const cookie = useCookie('auth-token', {
      maxAge: 60 * 60 * 24 * 7, // 7 dias
      path: '/',
    })
    cookie.value = response.token
  }

  function setUser(newUser: User) {
    user.value = newUser
  }

  function clearAuth() {
    token.value = null
    user.value = null
    
    // Limpar cookie
    const cookie = useCookie('auth-token')
    cookie.value = null
  }

  return {
    user,
    token,
    isAuthenticated,
    setAuth,
    setUser,
    clearAuth,
  }
}) 