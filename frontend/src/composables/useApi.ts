import { useAuthStore } from '@/stores/auth'
import { useRuntimeConfig } from '#app'

interface FetchOptions extends RequestInit {
  body?: any
}

export const useApi = () => {
  const authStore = useAuthStore()
  const config = useRuntimeConfig()
  const baseUrl = config.public.apiBase || 'http://localhost:8080/api'

  const fetch = async <T>(endpoint: string, options: FetchOptions = {}): Promise<T> => {
    const { body, ...fetchOptions } = options

    const headers = new Headers({
      'Content-Type': 'application/json',
      ...(fetchOptions.headers as Record<string, string> || {}),
    })

    // Adiciona o token de autenticação se existir
    const token = authStore.token
    if (token) {
      headers.set('Authorization', `Bearer ${token}`)
    }

    const response = await window.fetch(`${baseUrl}${endpoint}`, {
      ...fetchOptions,
      headers,
      body: body ? JSON.stringify(body) : undefined,
    })

    const data = await response.json()

    if (!response.ok) {
      throw data
    }

    return data as T
  }

  return {
    fetch
  }
} 