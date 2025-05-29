# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

# Instalar dependências necessárias
RUN apt-get update && apt-get install -y git ca-certificates tzdata && \
    update-ca-certificates

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Download de dependências com cache
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copiar código fonte
COPY . .

# Compilar a aplicação com flags de segurança
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main ./cmd/api

# Development stage
FROM golang:1.23 AS dev

WORKDIR /app

# Instalar dependências necessárias
RUN apt-get update && apt-get install -y git ca-certificates tzdata && \
    update-ca-certificates

# Instalar air para hot reload
RUN go install github.com/cosmtrek/air@latest

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Download de dependências com cache
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copiar código fonte
COPY . .

# Expor porta
EXPOSE 8080

# Comando para executar a aplicação em desenvolvimento
CMD ["air", "-c", ".air.toml"]

# Final stage
FROM alpine:3.19

WORKDIR /app

# Criar usuário não-root
RUN adduser -D -g '' appuser

# Instalar certificados e timezone
RUN apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates

# Copiar binário, configurações e documentação
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/cmd/api/docs/* ./swagger/

# Configurar permissões
RUN chown -R appuser:appuser /app && \
    chmod -R 755 /app && \
    chmod -R 644 /app/swagger/*

USER appuser

# Expor porta
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Comando para executar a aplicação
CMD ["./main"] 