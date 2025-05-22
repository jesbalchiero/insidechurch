<template>
  <div class="form-control w-full">
    <label v-if="label" :for="id" class="label">
      <span class="label-text">{{ label }}</span>
    </label>
    <input
      :id="id"
      :type="type"
      :value="modelValue"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      :placeholder="placeholder"
      :disabled="disabled"
      class="input input-bordered w-full"
      :class="{ 'input-error': error }"
    />
    <label v-if="error" class="label">
      <span class="label-text-alt text-error">{{ error }}</span>
    </label>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  modelValue: string
  label?: string
  type?: string
  placeholder?: string
  error?: string
  disabled?: boolean
  id?: string
}>()

defineEmits<{
  'update:modelValue': [value: string]
}>()
</script> 