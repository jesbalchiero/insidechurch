<template>
  <div class="min-h-screen flex flex-col md:flex-row">
    <!-- Lado esquerdo: Formulário -->
    <div class="flex-1 flex items-center justify-center p-8 bg-base-100">
      <div class="w-full max-w-md">
        <slot />
      </div>
    </div>

    <!-- Lado direito: Banner -->
    <div class="hidden md:flex flex-1 bg-gradient-to-br from-primary to-secondary text-primary-content">
      <div class="flex flex-col justify-center p-12">
        <h1 class="text-4xl font-bold mb-8">InsideChurch</h1>
        
        <div class="space-y-6">
          <div class="flex items-center space-x-4">
            <div class="w-12 h-12 rounded-full bg-primary-content/20 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-semibold">Gestão Simplificada</h3>
              <p class="text-primary-content/80">Gerencie sua igreja de forma eficiente e organizada</p>
            </div>
          </div>

          <div class="flex items-center space-x-4">
            <div class="w-12 h-12 rounded-full bg-primary-content/20 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-semibold">Comunicação Eficiente</h3>
              <p class="text-primary-content/80">Mantenha todos os membros conectados e informados</p>
            </div>
          </div>

          <div class="flex items-center space-x-4">
            <div class="w-12 h-12 rounded-full bg-primary-content/20 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-semibold">Segurança Garantida</h3>
              <p class="text-primary-content/80">Seus dados protegidos com as melhores práticas</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
</script> 