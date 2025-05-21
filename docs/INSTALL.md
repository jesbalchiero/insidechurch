# Guia de Instalação do InsideChurch

Este guia fornece instruções detalhadas para instalar e configurar o InsideChurch em diferentes ambientes.

## Índice
1. [Requisitos do Sistema](#requisitos-do-sistema)
2. [Instalação com Docker](#instalação-com-docker)
3. [Instalação Manual](#instalação-manual)
4. [Configuração do Banco de Dados](#configuração-do-banco-de-dados)
5. [Configuração do Ambiente](#configuração-do-ambiente)
6. [Solução de Problemas](#solução-de-problemas)

## Requisitos do Sistema

### Para Instalação com Docker
- Docker 20.10+
- Docker Compose 2.0+
- 2GB RAM mínimo
- 10GB espaço em disco

### Para Instalação Manual
- Go 1.21+
- PostgreSQL 15+
- Git
- 2GB RAM mínimo
- 10GB espaço em disco

## Instalação com Docker

### 1. Clone o Repositório
```bash
git clone https://github.com/seu-usuario/insidechurch.git
cd insidechurch
```

### 2. Configure as Variáveis de Ambiente
```bash
cp .env.example .env
```

Edite o arquivo `.env` com suas configurações:
```env
# Banco de dados
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha_segura
DB_NAME=insidechurch
DB_SSLMODE=disable

# JWT
JWT_SECRET=sua_chave_secreta_muito_segura

# Porta do servidor
PORT=8080
```

### 3. Inicie os Containers
```bash
docker-compose up -d
```

### 4. Verifique os Logs
```bash
docker-compose logs -f
```

## Instalação Manual

### 1. Instale o Go
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# macOS
brew install go

# Windows
# Baixe o instalador em https://golang.org/dl/
```

### 2. Instale o PostgreSQL
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install postgresql postgresql-contrib

# macOS
brew install postgresql

# Windows
# Baixe o instalador em https://www.postgresql.org/download/windows/
```

### 3. Configure o Banco de Dados
```sql
CREATE DATABASE insidechurch;
CREATE USER insidechurch WITH PASSWORD 'sua_senha_segura';
GRANT ALL PRIVILEGES ON DATABASE insidechurch TO insidechurch;
```

### 4. Clone e Configure o Projeto
```bash
git clone https://github.com/seu-usuario/insidechurch.git
cd insidechurch
go mod download
cp .env.example .env
```

### 5. Execute a Aplicação
```bash
go run cmd/api/main.go
```

## Configuração do Banco de Dados

### Migrações
O projeto usa GORM para migrações automáticas. As tabelas são criadas automaticamente na primeira execução.

### Backup
```bash
# Backup
pg_dump -U postgres insidechurch > backup.sql

# Restauração
psql -U postgres insidechurch < backup.sql
```

## Configuração do Ambiente

### Variáveis de Ambiente
| Variável | Descrição | Padrão |
|----------|-----------|---------|
| DB_HOST | Host do banco de dados | localhost |
| DB_PORT | Porta do banco de dados | 5432 |
| DB_USER | Usuário do banco | postgres |
| DB_PASSWORD | Senha do banco | - |
| DB_NAME | Nome do banco | insidechurch |
| DB_SSLMODE | Modo SSL | disable |
| JWT_SECRET | Chave secreta JWT | - |
| PORT | Porta da API | 8080 |

### Logs
Os logs são configurados para:
- Nível: INFO
- Formato: Texto com timestamp
- Saída: stdout

## Solução de Problemas

### Problemas Comuns

1. **Erro de Conexão com o Banco**
   - Verifique se o PostgreSQL está rodando
   - Confirme as credenciais no `.env`
   - Teste a conexão: `psql -U postgres -h localhost`

2. **Erro de Porta em Uso**
   - Verifique se a porta 8080 está livre
   - Mude a porta no `.env`

3. **Erro de Permissão**
   - Verifique as permissões do diretório
   - Execute como sudo se necessário

### Logs de Erro
Os logs de erro são salvos em:
- Docker: `docker-compose logs`
- Manual: stdout

### Suporte
Para mais ajuda:
- Abra uma issue no GitHub
- Consulte a documentação
- Entre em contato com o suporte 