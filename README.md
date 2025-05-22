# InsideChurch

Sistema completo para gerenciamento de igrejas e membros, desenvolvido com Go no backend e React no frontend.

## 🚀 Tecnologias

### Backend
- Go 1.21
- Gin (Framework Web)
- GORM (ORM)
- PostgreSQL
- JWT (Autenticação)
- Docker
- Swagger (Documentação da API)

### Frontend
- React
- TypeScript
- Material-UI
- Docker

## 📋 Pré-requisitos

- Go 1.21 ou superior
- Node.js 18 ou superior
- PostgreSQL 15
- Docker e Docker Compose

## 🔧 Instalação

### Usando Docker (Recomendado)

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/insidechurch.git
cd insidechurch
```

2. Configure as variáveis de ambiente:
```bash
# Backend
cp backend/.env.example backend/.env
# Frontend
cp frontend/.env.example frontend/.env
# Edite os arquivos .env com suas configurações
```

3. Inicie os containers:
```bash
docker-compose up -d
```

### Instalação Manual

#### Backend

1. Entre na pasta do backend:
```bash
cd backend
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

4. Execute a aplicação:
```bash
go run cmd/api/main.go
```

#### Frontend

1. Entre na pasta do frontend:
```bash
cd frontend
```

2. Instale as dependências:
```bash
npm install
```

3. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

4. Execute a aplicação:
```bash
npm start
```

## 🏗️ Arquitetura

O projeto está organizado em uma estrutura monorepo:

```
.
├── backend/           # API REST em Go
│   ├── cmd/
│   │   └── api/      # Ponto de entrada da aplicação
│   ├── internal/
│   │   ├── core/     # Regras de negócio e entidades
│   │   ├── handlers/ # Controladores HTTP
│   │   └── ...       # Outros componentes
│   └── tests/        # Testes de integração
├── frontend/         # Interface em React
│   ├── src/
│   ├── public/
│   └── ...
└── docker-compose.yml # Orquestração dos serviços
```

## 📚 Documentação da API

A documentação completa da API está disponível via Swagger em:
```
http://localhost:8080/swagger/index.html
```

## 🔐 Autenticação

A API usa JWT para autenticação. Para acessar rotas protegidas:

1. Registre um usuário:
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"email":"usuario@exemplo.com","password":"senha123","name":"Usuário Teste"}'
```

2. Faça login:
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"usuario@exemplo.com","password":"senha123"}'
```

3. Use o token retornado nas requisições:
```bash
curl -X GET http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer seu-token-jwt"
```

## 🧪 Testes

### Backend

#### Testes Unitários
```bash
cd backend
go test ./internal/...
```

#### Testes de Integração
```bash
cd backend
go test ./tests/integration/...
```

### Frontend

```bash
cd frontend
npm test
```

## 📦 Deploy

### Docker
```bash
docker-compose up -d
```

### Kubernetes
Exemplo de deployment:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: insidechurch
spec:
  replicas: 3
  template:
    spec:
      containers:
      - name: insidechurch-backend
        image: insidechurch-backend:latest
        ports:
        - containerPort: 8080
      - name: insidechurch-frontend
        image: insidechurch-frontend:latest
        ports:
        - containerPort: 3000
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie sua branch (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

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
