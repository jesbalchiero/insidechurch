# InsideChurch

Sistema completo para gerenciamento de igrejas e membros, desenvolvido com Go no backend e Nuxt.js no frontend.

## 📁 Estrutura do Projeto

```
insidechurch/
├── backend/           # API Go
│   ├── cmd/          # Ponto de entrada
│   ├── internal/     # Código interno
│   └── pkg/          # Pacotes públicos
├── frontend/         # Interface Nuxt.js
│   ├── components/   # Componentes Vue
│   ├── pages/        # Páginas
│   └── stores/       # Stores Pinia
├── docs/             # Documentação
└── scripts/          # Scripts úteis
```

## 🚀 Tecnologias

### Backend
- Go 1.21+
- Gin (Framework Web)
- GORM (ORM)
- PostgreSQL
- JWT (Autenticação)
- Docker

### Frontend
- Nuxt 3
- Vue 3
- TypeScript
- Tailwind CSS
- Pinia
- Vue Test Utils

## 🛠️ Instalação

### Usando Docker (Recomendado)

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/insidechurch.git
cd insidechurch
```

2. Configure as variáveis de ambiente:
```bash
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
```

3. Inicie os serviços:
```bash
./scripts/docker-dev.sh
```

### Instalação Manual

#### Backend

1. Instale o Go 1.21+
2. Configure o PostgreSQL
3. Instale dependências:
```bash
cd backend
go mod download
```

4. Execute:
```bash
go run cmd/api/main.go
```

#### Frontend

1. Instale Node.js 18+
2. Instale dependências:
```bash
cd frontend
npm install
```

3. Execute:
```bash
npm run dev
```

## 🏗️ Arquitetura

### Backend

Seguindo Clean Architecture:

- **Entities**: Regras de negócio
- **Use Cases**: Casos de uso
- **Interfaces**: Controllers e Repositories
- **Frameworks**: Gin, GORM, etc.

### Frontend

Arquitetura baseada em componentes:

- **Components**: Reutilizáveis e específicos
- **Pages**: Rotas da aplicação
- **Stores**: Gerenciamento de estado
- **Composables**: Lógica reutilizável

## 📚 Documentação

- [Uso da API](docs/API_USAGE.md)
- [Guia de Contribuição](CONTRIBUTING.md)
- [Checklist de Testes](docs/TESTS.md)

## 🧪 Testes

### Backend
```bash
cd backend
go test ./...
```

### Frontend
```bash
cd frontend
npm run test
npm run test:e2e
```

## 🚢 Deploy

### Docker

1. Build das imagens:
```bash
docker-compose -f docker-compose.prod.yml build
```

2. Iniciar serviços:
```bash
docker-compose -f docker-compose.prod.yml up -d
```

### Kubernetes

1. Aplique os manifests:
```bash
kubectl apply -f k8s/
```

2. Verifique o status:
```bash
kubectl get pods
```

## 📫 Suporte

- Abra uma issue para bugs
- Use discussions para ideias
- Entre em contato com mantenedores

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 📫 Contato

Seu Nome - [@seutwitter](https://twitter.com/seutwitter) - email@exemplo.com

Link do Projeto: [https://github.com/seu-usuario/insidechurch](https://github.com/seu-usuario/insidechurch)

## Ambiente de Desenvolvimento

### Requisitos
- Docker
- Docker Compose
- Node.js 18+
- Go 1.21+

### Portas
- Frontend: http://localhost:3000
- Backend: http://localhost:8080
- PostgreSQL: localhost:5432

### Comandos Docker

#### Desenvolvimento
Para iniciar o ambiente de desenvolvimento com Docker:

```bash
# Usar o script de desenvolvimento
./scripts/docker-dev.sh

# Ou manualmente
docker-compose down
docker-compose build --no-cache
docker-compose up
```

#### Produção
Para build de produção:

```bash
# Build das imagens
docker-compose -f docker-compose.prod.yml build

# Iniciar serviços
docker-compose -f docker-compose.prod.yml up -d
```

### Variáveis de Ambiente

#### Backend
- `DB_HOST`: Host do PostgreSQL (default: postgres)
- `DB_USER`: Usuário do PostgreSQL (default: postgres)
- `DB_PASSWORD`: Senha do PostgreSQL (default: postgres)
- `DB_NAME`: Nome do banco de dados (default: insidechurch)
- `DB_PORT`: Porta do PostgreSQL (default: 5432)
- `JWT_SECRET`: Chave secreta para JWT

#### Frontend
- `NUXT_PUBLIC_API_BASE`: URL base da API (default: http://localhost:8080)

### Estrutura do Projeto
```
.
├── backend/           # API Go
├── frontend/         # Frontend Nuxt.js
├── scripts/          # Scripts de desenvolvimento
└── docker-compose.yml
```
