<template>
  <div>
    <h1 class="text-3xl font-bold mb-8">
      Bem-vindo, {{ user?.value?.name }}!
    </h1>

    <!-- Cards de Estatísticas -->
    <div class="grid md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <!-- Total de Membros -->
      <div class="card bg-base-100 shadow-xl">
        <div class="card-body">
          <div class="flex items-center justify-between">
            <div>
              <h3 class="text-lg font-semibold">Total de Membros</h3>
              <p class="text-3xl font-bold">1.234</p>
            </div>
            <div class="w-12 h-12 rounded-full bg-primary/20 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-primary" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
          </div>
          <div class="text-sm text-base-content/60">
            +12% este mês
          </div>
        </div>
      </div>

      <!-- Presença Média -->
      <div class="card bg-base-100 shadow-xl">
        <div class="card-body">
          <div class="flex items-center justify-between">
            <div>
              <h3 class="text-lg font-semibold">Presença Média</h3>
              <p class="text-3xl font-bold">789</p>
            </div>
            <div class="w-12 h-12 rounded-full bg-secondary/20 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-secondary" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
              </svg>
            </div>
          </div>
          <div class="text-sm text-base-content/60">
            +5% este mês
          </div>
        </div>
      </div>

      <!-- Dízimos e Ofertas -->
      <div class="card bg-base-100 shadow-xl">
        <div class="card-body">
          <div class="flex items-center justify-between">
            <div>
              <h3 class="text-lg font-semibold">Dízimos e Ofertas</h3>
              <p class="text-3xl font-bold">R$ 45.678</p>
            </div>
            <div class="w-12 h-12 rounded-full bg-accent/20 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-accent" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </div>
          <div class="text-sm text-base-content/60">
            +8% este mês
          </div>
        </div>
      </div>

      <!-- Eventos -->
      <div class="card bg-base-100 shadow-xl">
        <div class="card-body">
          <div class="flex items-center justify-between">
            <div>
              <h3 class="text-lg font-semibold">Eventos</h3>
              <p class="text-3xl font-bold">12</p>
            </div>
            <div class="w-12 h-12 rounded-full bg-info/20 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-info" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
            </div>
          </div>
          <div class="text-sm text-base-content/60">
            3 eventos esta semana
          </div>
        </div>
      </div>
    </div>

    <!-- Próximos Eventos -->
    <div class="card bg-base-100 shadow-xl">
      <div class="card-body">
        <h2 class="card-title mb-4">Próximos Eventos</h2>
        <div class="space-y-4">
          <div v-for="i in 3" :key="i" class="flex items-center justify-between p-4 bg-base-200 rounded-lg">
            <div>
              <h3 class="font-semibold">Culto de Domingo</h3>
              <p class="text-sm text-base-content/60">Domingo, 10:00</p>
            </div>
            <button class="btn btn-sm">Ver Detalhes</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '~/composables/useAuth'

definePageMeta({
  layout: 'dashboard',
  middleware: ['auth']
})

const auth = useAuth()
const user = computed(() => auth.user)

onMounted(async () => {
  try {
    await auth.getUser()
  } catch (error) {
    console.error('Erro ao carregar usuário:', error)
  }
})
</script> 