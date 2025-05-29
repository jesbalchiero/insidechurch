# Inside Church Frontend

Frontend da aplicação Inside Church, desenvolvido com Nuxt 3, TypeScript, Tailwind CSS e Pinia.

## Requisitos

- Node.js 18.x ou superior
- npm 9.x ou superior

## Instalação

1. Clone o repositório
2. Instale as dependências:

```bash
npm install
```

3. Copie o arquivo de ambiente:

```bash
cp .env.example .env
```

4. Configure as variáveis de ambiente no arquivo `.env`

## Desenvolvimento

Para iniciar o servidor de desenvolvimento:

```bash
npm run dev
```

O servidor estará disponível em `http://localhost:3000`.

## Build

Para gerar a versão de produção:

```bash
npm run build
```

Para visualizar a versão de produção:

```bash
npm run preview
```

## Linting e Formatação

Para verificar o código:

```bash
npm run lint
```

Para corrigir problemas de linting:

```bash
npm run lint:fix
```

Para formatar o código:

```bash
npm run format
```

## Estrutura do Projeto

```
frontend/
├── components/     # Componentes Vue reutilizáveis
├── layouts/        # Layouts da aplicação
├── middleware/     # Middlewares de rota
├── pages/          # Páginas da aplicação
├── plugins/        # Plugins Nuxt
├── public/         # Arquivos estáticos
├── stores/         # Stores Pinia
└── types/          # Definições de tipos TypeScript
```

## Tecnologias Utilizadas

- [Nuxt 3](https://nuxt.com/)
- [TypeScript](https://www.typescriptlang.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Pinia](https://pinia.vuejs.org/)
- [Vue Router](https://router.vuejs.org/)
- [ESLint](https://eslint.org/)
- [Prettier](https://prettier.io/)
