import type { NitroFetchOptions } from 'nitropack'
import { useAuthStore } from '~/stores/auth'
import { useToast } from 'vue-toastification'

export const useApi = () => {
  const config = useRuntimeConfig()
  const router = useRouter()
  const auth = useAuthStore()
  const toast = useToast()

  const baseFetch = $fetch.create({
    baseURL: config.public.apiBase as string,
    headers: {
      'Content-Type': 'application/json',
    },
  })

  const handleError = (error: any) => {
    if (error.response?.status === 500) {
      toast.error('Ocorreu um erro interno. Por favor, tente novamente mais tarde.')
    }
    throw error
  }

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
      return handleError(error)
    }
  }

  return {
    fetch: fetchWithAuth,
  }
} 