import { defineNuxtConfig } from 'nuxt/config'

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt',
  ],
  imports: {
    dirs: ['stores', 'composables', 'types'],
  },
  css: [
    'vue-toastification/dist/index.css',
    '@/assets/css/main.css'
  ],
  plugins: [
    '~/plugins/toast.client.ts'
  ],
  typescript: {
    strict: true,
    typeCheck: true,
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
      apiBase: process.env.API_BASE || 'http://localhost:8080'
    }
  }
})
