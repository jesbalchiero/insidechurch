import type { NitroFetchOptions } from 'nitropack'
import { useAuthStore } from '~/stores/auth'

export const useApi = () => {
  const config = useRuntimeConfig()
  const router = useRouter()
  const auth = useAuthStore()

  const baseFetch = $fetch.create({
    baseURL: config.public.apiBase as string,
    headers: {
      'Content-Type': 'application/json',
    },
  })

  const fetchWithAuth = async <T>(
    url: string,
    options: {
      method?: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'
      body?: any
      headers?: Record<string, string>
    } = {}
  ) => {
    const headers: Record<string, string> = {
      ...(options.headers || {}),
      ...(auth.token ? { Authorization: `Bearer ${auth.token}` } : {}),
    }

    try {
      return await baseFetch<T>(url, {
        ...options,
        headers,
      })
    } catch (error: any) {
      if (error.statusCode === 401) {
        auth.clearAuth()
        router.push('/login')
      }
      throw error
    }
  }

  return {
    fetch: fetchWithAuth,
  }
} 