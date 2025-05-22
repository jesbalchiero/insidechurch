<template>
  <form @submit.prevent="handleSubmit" class="space-y-4">
    <BaseInput
      v-model="form.name"
      label="Nome"
      type="text"
      placeholder="Seu nome"
      :error="errors.name"
      :disabled="loading"
      required
    />

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

    <BaseInput
      v-model="form.confirmPassword"
      label="Confirmar Senha"
      type="password"
      placeholder="••••••••"
      :error="errors.confirmPassword"
      :disabled="loading"
      required
    />

    <BaseButton
      type="submit"
      variant="primary"
      :loading="loading"
      :disabled="loading"
    >
      Cadastrar
    </BaseButton>

    <div class="text-center">
      <NuxtLink to="/login" class="link link-primary">
        Já tem uma conta? Faça login
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
  name: '',
  email: '',
  password: '',
  confirmPassword: ''
})

const schema = z.object({
  name: z.string().min(1, 'Nome é obrigatório'),
  email: z.string().email('Email inválido'),
  password: z
    .string()
    .min(8, 'A senha deve ter no mínimo 8 caracteres')
    .regex(/[A-Z]/, 'A senha deve conter pelo menos uma letra maiúscula')
    .regex(/[a-z]/, 'A senha deve conter pelo menos uma letra minúscula')
    .regex(/[0-9]/, 'A senha deve conter pelo menos um número')
    .regex(/[!@#$%^&*]/, 'A senha deve conter pelo menos um caractere especial (!@#$%^&*)'),
  confirmPassword: z.string()
}).refine((data) => data.password === data.confirmPassword, {
  message: 'As senhas não coincidem',
  path: ['confirmPassword']
})

const handleSubmit = async () => {
  try {
    errors.value = {}
    const validatedData = schema.parse(form.value)
    loading.value = true
    await auth.register({
      name: validatedData.name,
      email: validatedData.email,
      password: validatedData.password
    })
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