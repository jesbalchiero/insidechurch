<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="max-w-md w-full p-6 bg-white rounded-lg shadow-lg">
      <h1 class="text-3xl font-bold text-red-600 mb-4">
        {{ error?.statusCode === 404 ? 'Página não encontrada' : 'Erro' }}
      </h1>
      <p class="text-gray-600 mb-6">
        {{ error?.message || 'Ocorreu um erro inesperado.' }}
      </p>
      <button
        @click="handleError"
        class="w-full bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700 transition-colors"
      >
        Voltar para a página inicial
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { clearError } from '#app'

interface ErrorProps {
  error: {
    statusCode?: number
    message?: string
  }
}

const props = defineProps<ErrorProps>()

const handleError = () => {
  clearError({ redirect: '/' })
}
</script> 