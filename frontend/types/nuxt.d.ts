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
    $fetch: typeof $fetch
  }

  interface RuntimeConfig {
    public: {
      apiBase: string
    }
  }

  export function defineNuxtPlugin(plugin: (nuxtApp: NuxtApp) => any): any
  export function useRuntimeConfig(): RuntimeConfig
  export function clearError(options?: { redirect?: string }): void
  export type { NuxtApp, RuntimeConfig }
}

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $fetch: typeof $fetch
  }
}

declare module 'vue' {
  interface ComponentCustomProperties {
    $fetch: typeof $fetch
  }
}

export {} 