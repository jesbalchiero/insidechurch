<template>
  <div class="min-h-screen flex items-center justify-center bg-base-200">
    <div class="card w-full max-w-md bg-base-100 shadow-xl">
      <div class="card-body">
        <h1 class="text-2xl font-bold text-center mb-6">Cadastro</h1>
        <RegisterForm />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuth } from '../composables/useAuth'
import { useRouter } from 'vue-router'
import { useToast } from '../composables/useToast'

const auth = useAuth()
const router = useRouter()
const toast = useToast()

const form = ref({
  name: '',
  email: '',
  password: '',
  confirmPassword: ''
})

const handleSubmit = async () => {
  try {
    await auth.register(form.value)
    toast.success('Conta criada com sucesso!')
    router.push('/login')
  } catch (error: any) {
    toast.error(error.response?._data?.message || 'Erro ao criar conta')
  }
}

definePageMeta({
  layout: 'auth',
  middleware: ['guest']
})
</script> 