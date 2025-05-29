import type { ToastInterface } from 'vue-toastification'

declare module '#app' {
  interface NuxtApp {
    $toast: ToastInterface
  }
}

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $toast: ToastInterface
  }
} 