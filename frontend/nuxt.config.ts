// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt',
  ],
  pinia: {
    storesDirs: ['./stores/**'],
  },
  imports: {
    dirs: ['stores'],
  },
  css: [
    'vue-toastification/dist/index.css',
    '@/assets/css/main.css'
  ],
  tailwindcss: {
    cssPath: '~/assets/css/main.css',
    configPath: 'tailwind.config.js',
    exposeConfig: false,
    viewer: true,
  },
  app: {
    head: {
      title: 'Inside Church',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Sistema de gerenciamento para igrejas' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ]
    }
  },
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080'
    }
  }
})
