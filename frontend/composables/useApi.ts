import { useRuntimeConfig } from 'nuxt/app'
import { useAuthStore } from '~/stores/auth'

export const useApi = () => {
  const config = useRuntimeConfig()
  const authStore = useAuthStore()

  const fetch = async <T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> => {
    const url = `${config.public.apiBase}${endpoint}`

    const headers = new Headers({
      'Content-Type': 'application/json',
      ...(options.headers as Record<string, string>),
    })

    if (authStore.token) {
      headers.set('Authorization', `Bearer ${authStore.token}`)
    }

    const response = await globalThis.fetch(url, {
      ...options,
      headers,
    })

    if (!response.ok) {
      const error = await response.text()
      throw new Error(error)
    }

    return response.json()
  }

  return {
    fetch,
  }
} 