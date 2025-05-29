<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Inside Church
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          Bem-vindo ao sistema de gerenciamento da sua igreja
        </p>
      </div>
      <div class="text-center space-y-4">
        <NuxtLink
          to="/login"
          class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Entrar
        </NuxtLink>
        <div>
          <NuxtLink
            to="/register"
            class="font-medium text-indigo-600 hover:text-indigo-500"
          >
            Criar uma nova conta
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'default',
  middleware: ['guest']
})
</script> 