import Toast, { type PluginOptions, POSITION } from 'vue-toastification'
import 'vue-toastification/dist/index.css'
import { defineNuxtPlugin } from '#app'
import type { NuxtApp } from '#app'

export default defineNuxtPlugin((nuxtApp: NuxtApp) => {
  const options: PluginOptions = {
    position: POSITION.TOP_RIGHT,
    timeout: 5000,
    closeOnClick: true,
    pauseOnFocusLoss: true,
    pauseOnHover: true,
    draggable: true,
    draggablePercent: 0.6,
    showCloseButtonOnHover: false,
    hideProgressBar: false,
    closeButton: 'button',
    icon: true,
    rtl: false
  }

  nuxtApp.vueApp.use(Toast, options)

  return {
    provide: {
      toast: nuxtApp.vueApp.config.globalProperties.$toast
    }
  }
}) 