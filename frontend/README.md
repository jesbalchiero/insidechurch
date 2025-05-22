# Frontend - InsideChurch

Interface do usuário desenvolvida com Nuxt 3, Vue 3 e TypeScript.

## 🚀 Stack Tecnológica

- **Framework**: Nuxt 3
- **UI**: Vue 3 + Composition API
- **Estilização**: Tailwind CSS
- **Estado**: Pinia
- **Tipagem**: TypeScript
- **Testes**: Vitest + Vue Test Utils
- **E2E**: Cypress
- **Linting**: ESLint + Prettier
- **Notificações**: Vue Toastification

## 📁 Estrutura de Pastas

```
frontend/
├── components/        # Componentes Vue
│   ├── common/       # Componentes base reutilizáveis
│   └── features/     # Componentes específicos de features
├── composables/      # Lógica reutilizável
├── layouts/          # Layouts da aplicação
├── middleware/       # Middleware de rotas
├── pages/           # Páginas da aplicação
├── plugins/         # Plugins Vue
├── public/          # Arquivos estáticos
├── stores/          # Stores Pinia
├── types/           # Definições TypeScript
└── utils/           # Funções utilitárias
```

## 🎨 Padrões de Código

### Componentes

1. **Nomenclatura**:
   - PascalCase para componentes
   - kebab-case para arquivos
   - Sufixo `.vue` para componentes

2. **Estrutura**:
```vue
<template>
  <!-- Template -->
</template>

<script setup lang="ts">
// Imports
// Props
// Emits
// Composables
// Computed
// Methods
</script>

<style scoped>
/* Estilos */
</style>
```

3. **Props**:
```typescript
const props = defineProps<{
  title: string
  items: Item[]
  loading?: boolean
}>()
```

4. **Emits**:
```typescript
const emit = defineEmits<{
  (e: 'update', value: string): void
  (e: 'delete', id: string): void
}>()
```

### Composables

1. **Nomenclatura**:
   - Prefixo `use` (ex: `useAuth`, `useApi`)
   - Um arquivo por composable

2. **Estrutura**:
```typescript
export const useExample = () => {
  // Estado
  const state = ref()

  // Métodos
  const method = () => {}

  // Retorno
  return {
    state,
    method
  }
}
```

### Stores

1. **Nomenclatura**:
   - Sufixo `Store` (ex: `authStore`, `userStore`)
   - Um arquivo por store

2. **Estrutura**:
```typescript
export const useExampleStore = defineStore('example', {
  state: () => ({}),
  getters: {},
  actions: {}
})
```

## 🛠️ Como Adicionar Novos Componentes

1. **Componente Base**:
```bash
# Criar em components/common/
touch components/common/BaseExample.vue
```

2. **Componente de Feature**:
```bash
# Criar em components/features/
touch components/features/ExampleFeature.vue
```

3. **Composable**:
```bash
# Criar em composables/
touch composables/useExample.ts
```

4. **Store**:
```bash
# Criar em stores/
touch stores/exampleStore.ts
```

## 📦 Scripts Disponíveis

```bash
# Desenvolvimento
npm run dev

# Build
npm run build

# Preview
npm run preview

# Lint
npm run lint

# Testes
npm run test
npm run test:e2e
```

## 🔧 Variáveis de Ambiente

Crie um arquivo `.env` baseado no `.env.example`:

```env
NUXT_PUBLIC_API_BASE=http://localhost:8080
```

## 🧪 Testes

### Unitários
```bash
npm run test
```

### E2E
```bash
npm run test:e2e
```

## 📚 Documentação Adicional

- [Vue 3](https://vuejs.org/)
- [Nuxt 3](https://nuxt.com/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Pinia](https://pinia.vuejs.org/)
- [TypeScript](https://www.typescriptlang.org/)
