import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

export default defineNuxtPlugin((nuxtApp) => {
  const pinia = createPinia()
  
  // Configurar auto-imports
  pinia.use(({ store }) => {
    store.$subscribe((mutation, state) => {
      // Persistir estado automaticamente
      localStorage.setItem(store.$id, JSON.stringify(state))
    })
  })
  
  // Usar plugin de persistência
  pinia.use(piniaPluginPersistedstate)
  
  nuxtApp.vueApp.use(pinia)
}) 