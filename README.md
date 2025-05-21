# InsideChurch API

API REST para gerenciamento de igrejas e membros, desenvolvida em Go.

## 🚀 Tecnologias

- Go 1.21
- Gin (Framework Web)
- GORM (ORM)
- PostgreSQL
- JWT (Autenticação)
- Docker
- Swagger (Documentação da API)

## 📋 Pré-requisitos

- Go 1.21 ou superior
- PostgreSQL 15
- Docker e Docker Compose (opcional)

## 🔧 Instalação

### Usando Docker (Recomendado)

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/insidechurch.git
cd insidechurch
```

2. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

3. Inicie os containers:
```bash
docker-compose up -d
```

### Instalação Manual

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/insidechurch.git
cd insidechurch
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

## 🏗️ Arquitetura

O projeto segue uma arquitetura limpa (Clean Architecture) com as seguintes camadas:

```
.
├── cmd/
│   └── api/            # Ponto de entrada da aplicação
├── internal/
│   ├── core/           # Regras de negócio e entidades
│   │   ├── domain/     # Entidades e interfaces
│   │   └── interfaces/ # Interfaces dos repositórios
│   ├── handlers/       # Controladores HTTP
│   ├── middleware/     # Middlewares (auth, logging)
│   ├── repositories/   # Implementação dos repositórios
│   ├── routes/         # Configuração das rotas
│   └── services/       # Lógica de negócio
└── tests/              # Testes de integração
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

### Testes Unitários
```bash
go test ./internal/...
```

### Testes de Integração
```bash
go test ./tests/integration/...
```

## 📦 Deploy

### Docker
```bash
docker build -t insidechurch .
docker run -p 8080:8080 insidechurch
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
      - name: insidechurch
        image: insidechurch:latest
        ports:
        - containerPort: 8080
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
