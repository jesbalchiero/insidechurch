# Arquitetura do InsideChurch

Este documento descreve a arquitetura do sistema InsideChurch, suas camadas, componentes e decisões de design.

## Visão Geral

O InsideChurch segue os princípios da Clean Architecture, separando o código em camadas bem definidas e mantendo as regras de negócio independentes de frameworks e detalhes externos.

## Camadas da Aplicação

### 1. Core (Domínio)

Localização: `internal/core/`

#### Domain
- Contém as entidades do negócio
- Define as estruturas de dados principais
- Independente de frameworks

Exemplo:
```go
type User struct {
    gorm.Model
    Email    string
    Password string
    Name     string
}
```

#### Interfaces
- Define contratos para repositórios
- Permite inversão de dependência
- Facilita testes unitários

Exemplo:
```go
type UserRepository interface {
    Create(user *domain.User) error
    FindByEmail(email string) (*domain.User, error)
    FindByID(id uint) (*domain.User, error)
}
```

### 2. Infraestrutura

Localização: `internal/`

#### Repositories
- Implementa as interfaces do domínio
- Gerencia persistência de dados
- Usa GORM para operações no banco

#### Services
- Implementa regras de negócio
- Coordena operações entre repositórios
- Gerencia autenticação e autorização

#### Handlers
- Gerencia requisições HTTP
- Valida dados de entrada
- Formata respostas

#### Middleware
- Gerencia autenticação
- Logging
- Tratamento de erros

#### Routes
- Define endpoints da API
- Configura middlewares
- Organiza rotas por domínio

## Fluxo de Dados

1. **Requisição HTTP**
   ```
   Cliente -> Router -> Middleware -> Handler
   ```

2. **Processamento**
   ```
   Handler -> Service -> Repository -> Database
   ```

3. **Resposta**
   ```
   Database -> Repository -> Service -> Handler -> Cliente
   ```

## Padrões de Design

### 1. Repository Pattern
- Abstrai acesso a dados
- Facilita troca de banco de dados
- Melhora testabilidade

### 2. Dependency Injection
- Injeção via construtores
- Facilita testes unitários
- Reduz acoplamento

### 3. Middleware Chain
- Processamento em pipeline
- Reutilização de código
- Separação de responsabilidades

## Segurança

### 1. Autenticação
- JWT para tokens
- Senhas hasheadas com bcrypt
- Tokens com expiração

### 2. Autorização
- Middleware de autenticação
- Validação de permissões
- Proteção de rotas

### 3. Validação
- Validação de entrada
- Sanitização de dados
- Prevenção de SQL Injection

## Testes

### 1. Unitários
- Testes de serviços
- Testes de repositórios
- Mocks e stubs

### 2. Integração
- Testes de API
- Testes de banco
- Testes de autenticação

## Logging e Monitoramento

### 1. Logs
- Logrus para logging estruturado
- Diferentes níveis de log
- Formato JSON em produção

### 2. Métricas
- Prometheus para métricas
- Health checks
- Monitoramento de performance

## Escalabilidade

### 1. Horizontal
- Stateless design
- Load balancing
- Cache distribuído

### 2. Vertical
- Otimização de queries
- Índices apropriados
- Connection pooling

## Deploy

### 1. Containerização
- Docker multi-stage
- Imagens otimizadas
- Variáveis de ambiente

### 2. Orquestração
- Kubernetes ready
- Health checks
- Auto-scaling

## Próximos Passos

1. **Cache**
   - Implementar Redis
   - Cache de queries
   - Cache de sessões

2. **Message Queue**
   - Processamento assíncrono
   - Eventos de domínio
   - Background jobs

3. **API Gateway**
   - Rate limiting
   - Circuit breaker
   - API versioning 