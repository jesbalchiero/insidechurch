import type { NitroFetchOptions } from 'nitropack'

export const useApi = () => {
  const config = useRuntimeConfig()
  const router = useRouter()

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
    const token = localStorage.getItem('token')

    const headers: Record<string, string> = {
      ...(options.headers || {}),
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    }

    try {
      return await baseFetch<T>(url, {
        ...options,
        headers,
      })
    } catch (error: any) {
      if (error.statusCode === 401) {
        localStorage.removeItem('token')
        router.push('/login')
      }
      throw error
    }
  }

  return {
    fetch: fetchWithAuth,
  }
} 