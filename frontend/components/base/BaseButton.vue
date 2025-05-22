<template>
  <button
    :type="type"
    :disabled="disabled || loading"
    class="btn"
    :class="[
      variant === 'primary' && 'btn-primary',
      variant === 'secondary' && 'btn-secondary',
      variant === 'accent' && 'btn-accent',
      variant === 'ghost' && 'btn-ghost',
      variant === 'link' && 'btn-link',
      'w-full'
    ]"
  >
    <span v-if="loading" class="loading loading-spinner loading-sm"></span>
    <slot v-else></slot>
  </button>
</template>

<script setup lang="ts">
defineProps<{
  type?: 'button' | 'submit' | 'reset'
  variant?: 'primary' | 'secondary' | 'accent' | 'ghost' | 'link'
  disabled?: boolean
  loading?: boolean
}>()
</script> 