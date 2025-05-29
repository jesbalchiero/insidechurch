declare function definePageMeta(meta: {
  layout?: string
  middleware?: string | string[]
  [key: string]: any
}): void 

declare module '#app' {
  interface NuxtApp {
    vueApp: {
      use: (plugin: any, options?: any) => void
      config: {
        globalProperties: {
          $toast: any
        }
      }
    }
  }

  interface RuntimeConfig {
    public: {
      apiBase: string
    }
  }

  export function defineNuxtPlugin(plugin: (nuxtApp: NuxtApp) => any): any
  export function useRuntimeConfig(): RuntimeConfig
  export type { NuxtApp, RuntimeConfig }
} 