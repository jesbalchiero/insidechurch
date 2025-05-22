<template>
  <div class="min-h-screen flex items-center justify-center bg-base-200">
    <div class="card w-full max-w-md bg-base-100 shadow-xl">
      <div class="card-body">
        <h1 class="text-2xl font-bold text-center mb-6">Login</h1>
        <LoginForm />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuth } from '../composables/useAuth'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'

const auth = useAuth()
const router = useRouter()
const toast = useToast()

const form = ref({
  email: '',
  password: ''
})

const handleSubmit = async () => {
  try {
    await auth.login(form.value)
    toast.success('Login realizado com sucesso!')
    router.push('/dashboard')
  } catch (error: any) {
    toast.error(error.response?._data?.message || 'Erro ao fazer login')
  }
}

definePageMeta({
  layout: 'auth',
  middleware: ['guest']
})
</script> 