import type { ToastInterface } from 'vue-toastification'

export const useToast = () => {
  const nuxtApp = useNuxtApp()
  return nuxtApp.$toast
} 