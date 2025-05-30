#!/bin/bash

# Cores para output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

# Função para verificar se um comando existe
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Verificar se o Docker está instalado
if ! command_exists docker; then
    echo -e "${RED}Docker não está instalado. Por favor, instale o Docker primeiro.${NC}"
    exit 1
fi

# Verificar se o Docker Compose está instalado
if ! command_exists docker-compose; then
    echo -e "${RED}Docker Compose não está instalado. Por favor, instale o Docker Compose primeiro.${NC}"
    exit 1
fi

# Parar containers existentes
echo -e "${YELLOW}Parando containers existentes...${NC}"
docker-compose down

# Remover volumes antigos
echo -e "${YELLOW}Removendo volumes antigos...${NC}"
docker volume rm insidechurch_postgres_data insidechurch_backend_cache insidechurch_frontend_cache 2>/dev/null || true

# Rebuildar imagens
echo -e "${YELLOW}Rebuildando imagens...${NC}"
docker-compose build --no-cache

# Iniciar containers
echo -e "${YELLOW}Iniciando containers...${NC}"
docker-compose up --force-recreate

# Aguardar todos os serviços estarem prontos
echo -e "${YELLOW}Aguardando serviços iniciarem...${NC}"
sleep 10

# Verificar status dos containers
echo -e "${YELLOW}Verificando status dos containers...${NC}"
if docker-compose ps | grep -q "Exit"; then
    echo -e "${RED}Alguns containers falharam ao iniciar. Verifique os logs:${NC}"
    docker-compose logs
    exit 1
fi

echo -e "${GREEN}Serviços disponíveis em:${NC}"
echo -e "Frontend: ${GREEN}http://localhost:3000${NC}"
echo -e "Backend:  ${GREEN}http://localhost:8080${NC}"
echo -e "Postgres: ${GREEN}localhost:5432${NC}" 