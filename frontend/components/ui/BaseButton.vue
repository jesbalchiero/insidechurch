<script setup lang="ts">
import LoadingSpinner from '~/components/common/LoadingSpinner.vue'

interface Props {
  variant?: 'primary' | 'secondary'
  size?: 'sm' | 'md' | 'lg'
  loading?: boolean
  disabled?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  size: 'md',
  loading: false,
  disabled: false,
})

const sizeClasses = {
  sm: 'px-2.5 py-1.5 text-xs',
  md: 'px-4 py-2 text-sm',
  lg: 'px-6 py-3 text-base',
}

const variantClasses = {
  primary: 'btn-primary',
  secondary: 'btn-secondary',
}
</script>

<template>
  <button
    :class="[
      variantClasses[variant],
      sizeClasses[size],
    ]"
    :disabled="disabled || loading"
  >
    <LoadingSpinner
      v-if="loading"
      size="sm"
      class="mr-2"
    />
    <slot />
  </button>
</template> 