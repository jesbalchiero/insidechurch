#!/bin/bash

# Iniciar o backend em background
cd backend && go run main.go &
BACKEND_PID=$!

# Iniciar o frontend em background
cd ../frontend && yarn dev &
FRONTEND_PID=$!

# Função para limpar os processos ao sair
cleanup() {
    echo "Encerrando serviços..."
    kill $BACKEND_PID
    kill $FRONTEND_PID
    exit 0
}

# Capturar CTRL+C
trap cleanup SIGINT

# Manter o script rodando
echo "Serviços iniciados! Pressione CTRL+C para encerrar."
wait 