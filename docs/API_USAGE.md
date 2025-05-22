# Guia de Uso da API InsideChurch

Este guia fornece exemplos práticos de como usar a API do InsideChurch.

## Autenticação

### Registro de Usuário

```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "usuario@exemplo.com",
    "password": "senha123",
    "name": "Usuário Teste"
  }'
```

Resposta:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "usuario@exemplo.com",
    "name": "Usuário Teste"
  }
}
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "usuario@exemplo.com",
    "password": "senha123"
  }'
```

Resposta:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "usuario@exemplo.com",
    "name": "Usuário Teste"
  }
}
```

## Usuários

### Obter Dados do Usuário Atual

```bash
curl -X GET http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer seu-token-jwt"
```

Resposta:
```json
{
  "id": 1,
  "email": "usuario@exemplo.com",
  "name": "Usuário Teste"
}
```

### Atualizar Dados do Usuário

```bash
curl -X PUT http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer seu-token-jwt" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Novo Nome"
  }'
```

Resposta:
```json
{
  "id": 1,
  "email": "usuario@exemplo.com",
  "name": "Novo Nome"
}
```

## Exemplos em Diferentes Linguagens

### Python

```python
import requests

BASE_URL = "http://localhost:8080/api/v1"
TOKEN = "seu-token-jwt"

# Headers com autenticação
headers = {
    "Authorization": f"Bearer {TOKEN}",
    "Content-Type": "application/json"
}

# Login
def login(email, password):
    response = requests.post(
        f"{BASE_URL}/users/login",
        json={"email": email, "password": password}
    )
    return response.json()

# Obter dados do usuário
def get_user_data():
    response = requests.get(
        f"{BASE_URL}/users/me",
        headers=headers
    )
    return response.json()
```

### JavaScript (Node.js)

```javascript
const axios = require('axios');

const BASE_URL = 'http://localhost:8080/api/v1';
const TOKEN = 'seu-token-jwt';

// Configuração do axios
const api = axios.create({
    baseURL: BASE_URL,
    headers: {
        'Authorization': `Bearer ${TOKEN}`,
        'Content-Type': 'application/json'
    }
});

// Login
async function login(email, password) {
    const response = await api.post('/users/login', {
        email,
        password
    });
    return response.data;
}

// Obter dados do usuário
async function getUserData() {
    const response = await api.get('/users/me');
    return response.data;
}
```

### Go

```go
package main

import (
    "fmt"
    "net/http"
    "bytes"
    "encoding/json"
)

const (
    baseURL = "http://localhost:8080/api/v1"
    token   = "seu-token-jwt"
)

// Login
func login(email, password string) (map[string]interface{}, error) {
    data := map[string]string{
        "email":    email,
        "password": password,
    }
    
    jsonData, _ := json.Marshal(data)
    
    resp, err := http.Post(
        baseURL+"/users/login",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)
    return result, nil
}

// Obter dados do usuário
func getUserData() (map[string]interface{}, error) {
    req, _ := http.NewRequest("GET", baseURL+"/users/me", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)
    return result, nil
}
```

## Tratamento de Erros

### Exemplos de Respostas de Erro

#### 400 Bad Request
```json
{
  "error": "Dados inválidos",
  "details": {
    "email": "Email inválido",
    "password": "Senha muito curta"
  }
}
```

#### 401 Unauthorized
```json
{
  "error": "Token inválido ou expirado"
}
```

#### 404 Not Found
```json
{
  "error": "Recurso não encontrado"
}
```

#### 500 Internal Server Error
```json
{
  "error": "Erro interno do servidor"
}
```

## Boas Práticas

1. **Autenticação**
   - Sempre use HTTPS em produção
   - Armazene tokens de forma segura
   - Implemente refresh token

2. **Requisições**
   - Use paginação para listas grandes
   - Implemente cache quando possível
   - Trate timeouts adequadamente

3. **Respostas**
   - Use códigos HTTP apropriados
   - Padronize formato de resposta
   - Inclua mensagens de erro claras

## Rate Limiting

A API implementa rate limiting para prevenir abusos:

- 100 requisições por minuto por IP
- 1000 requisições por hora por usuário

Headers de resposta:
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 99
X-RateLimit-Reset: 1612345678
```

## Webhooks

Para receber notificações de eventos:

1. Configure o webhook:
```bash
curl -X POST http://localhost:8080/api/v1/webhooks \
  -H "Authorization: Bearer seu-token-jwt" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://seu-servidor.com/webhook",
    "events": ["user.created", "user.updated"]
  }'
```

2. Receba eventos:
```json
{
  "event": "user.created",
  "data": {
    "id": 1,
    "email": "usuario@exemplo.com"
  },
  "timestamp": "2024-02-06T12:00:00Z"
}
```

# Documentação da API

## Backend

### Autenticação

#### Login
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "usuario@exemplo.com",
  "password": "senha123"
}
```

Resposta:
```json
{
  "token": "jwt_token",
  "user": {
    "id": 1,
    "name": "Usuário",
    "email": "usuario@exemplo.com"
  }
}
```

#### Registro
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "name": "Novo Usuário",
  "email": "novo@exemplo.com",
  "password": "senha123"
}
```

Resposta:
```json
{
  "token": "jwt_token",
  "user": {
    "id": 2,
    "name": "Novo Usuário",
    "email": "novo@exemplo.com"
  }
}
```

### Usuários

#### Listar Usuários
```http
GET /api/v1/users
Authorization: Bearer jwt_token
```

#### Obter Usuário
```http
GET /api/v1/users/:id
Authorization: Bearer jwt_token
```

#### Atualizar Usuário
```http
PUT /api/v1/users/:id
Authorization: Bearer jwt_token
Content-Type: application/json

{
  "name": "Nome Atualizado",
  "email": "atualizado@exemplo.com"
}
```

#### Deletar Usuário
```http
DELETE /api/v1/users/:id
Authorization: Bearer jwt_token
```

## Frontend

### Fluxo de Autenticação

1. **Login**
   - Usuário acessa `/login`
   - Preenche formulário com email e senha
   - Frontend valida campos
   - Envia requisição para `/api/v1/auth/login`
   - Armazena token e dados do usuário
   - Redireciona para dashboard

2. **Registro**
   - Usuário acessa `/register`
   - Preenche formulário com nome, email e senha
   - Frontend valida campos
   - Envia requisição para `/api/v1/auth/register`
   - Armazena token e dados do usuário
   - Redireciona para dashboard

3. **Logout**
   - Usuário clica em logout
   - Frontend remove token e dados do usuário
   - Redireciona para login

### Middleware de Autenticação

O frontend implementa dois middlewares:

1. **auth**
   - Protege rotas que requerem autenticação
   - Verifica token no localStorage
   - Redireciona para login se não autenticado

2. **guest**
   - Protege rotas de visitantes
   - Redireciona para dashboard se autenticado

### Plugin de Autenticação

- Executa apenas no cliente
- Verifica token ao iniciar aplicação
- Carrega dados do usuário se token existir
- Remove token se inválido

### Composables

1. **useAuth**
```typescript
const { user, login, register, logout } = useAuth()
```

2. **useApi**
```typescript
const { get, post, put, del } = useApi()
```

### Stores

1. **authStore**
```typescript
const { user, token, isAuthenticated } = useAuthStore()
```

2. **userStore**
```typescript
const { users, fetchUsers, updateUser } = useUserStore()
```

## Exemplos de Uso

### Login
```typescript
const { login } = useAuth()

try {
  await login({
    email: 'usuario@exemplo.com',
    password: 'senha123'
  })
  // Redirecionar para dashboard
} catch (error) {
  // Mostrar erro
}
```

### Listar Usuários
```typescript
const { users, fetchUsers } = useUserStore()

onMounted(async () => {
  await fetchUsers()
})
```

### Atualizar Usuário
```typescript
const { updateUser } = useUserStore()

try {
  await updateUser(id, {
    name: 'Novo Nome',
    email: 'novo@email.com'
  })
  // Mostrar sucesso
} catch (error) {
  // Mostrar erro
}
``` 