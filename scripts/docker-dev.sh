#!/bin/bash

# Cores para output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${YELLOW}Parando containers existentes...${NC}"
docker-compose down

echo -e "${YELLOW}Removendo volumes antigos...${NC}"
docker volume rm insidechurch_postgres_data || true

echo -e "${YELLOW}Rebuildando imagens...${NC}"
docker-compose build --no-cache

echo -e "${YELLOW}Iniciando containers...${NC}"
docker-compose up --force-recreate

echo -e "${GREEN}Serviços disponíveis em:${NC}"
echo -e "Frontend: ${GREEN}http://localhost:3000${NC}"
echo -e "Backend:  ${GREEN}http://localhost:8080${NC}"
echo -e "Postgres: ${GREEN}localhost:5432${NC}" 