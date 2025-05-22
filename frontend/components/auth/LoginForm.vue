<template>
  <form @submit.prevent="handleSubmit" class="space-y-4">
    <BaseInput
      v-model="form.email"
      label="Email"
      type="email"
      placeholder="seu@email.com"
      :error="errors.email"
      :disabled="loading"
      required
    />

    <BaseInput
      v-model="form.password"
      label="Senha"
      type="password"
      placeholder="••••••••"
      :error="errors.password"
      :disabled="loading"
      required
    />

    <BaseButton
      type="submit"
      variant="primary"
      :loading="loading"
      :disabled="loading"
    >
      Entrar
    </BaseButton>

    <div class="text-center">
      <NuxtLink to="/register" class="link link-primary">
        Não tem uma conta? Cadastre-se
      </NuxtLink>
    </div>
  </form>
</template>

<script setup lang="ts">
import { z } from 'zod'
import { useAuth } from '~/composables/useAuth'

const auth = useAuth()
const loading = ref(false)
const errors = ref<Record<string, string>>({})

const form = ref({
  email: '',
  password: ''
})

const schema = z.object({
  email: z.string().email('Email inválido'),
  password: z.string().min(1, 'Senha é obrigatória')
})

const handleSubmit = async () => {
  try {
    errors.value = {}
    const validatedData = schema.parse(form.value)
    loading.value = true
    await auth.login(validatedData)
  } catch (error) {
    if (error instanceof z.ZodError) {
      errors.value = error.errors.reduce((acc, err) => {
        if (err.path[0]) {
          acc[err.path[0] as string] = err.message
        }
        return acc
      }, {} as Record<string, string>)
    }
  } finally {
    loading.value = false
  }
}
</script> 