# Etapa 1: build
FROM golang:1.21.8-alpine3.19 AS builder
WORKDIR /app

# Instalar dependências de segurança e build
RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    git

# Copiar e baixar dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar a aplicação com flags de segurança
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -extldflags=-static" \
    -o insidechurch ./cmd/api

# Etapa 2: imagem final
FROM alpine:3.19.4
WORKDIR /app

# Instalar dependências de segurança
RUN apk add --no-cache \
    ca-certificates \
    tzdata

# Criar usuário não-root
RUN adduser -D -u 1000 appuser

# Copiar binário e configurações
COPY --from=builder /app/insidechurch .
COPY .env.example .env

# Configurar permissões
RUN chown -R appuser:appuser /app

# Mudar para usuário não-root
USER appuser

# Configurar variáveis de ambiente
ENV TZ=UTC

EXPOSE 8080
CMD ["./insidechurch"] 