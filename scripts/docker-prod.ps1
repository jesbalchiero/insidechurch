# Cores para output
$Green = [System.ConsoleColor]::Green
$Yellow = [System.ConsoleColor]::Yellow
$Red = [System.ConsoleColor]::Red

# Função para verificar se um comando existe
function Test-Command {
    param ($command)
    $oldPreference = $ErrorActionPreference
    $ErrorActionPreference = 'stop'
    try { if (Get-Command $command) { return $true } }
    catch { return $false }
    finally { $ErrorActionPreference = $oldPreference }
}

# Verificar se o Docker está instalado
if (-not (Test-Command docker)) {
    Write-Host "Docker não está instalado. Por favor, instale o Docker primeiro." -ForegroundColor $Red
    exit 1
}

# Verificar se o Docker Compose está instalado
if (-not (Test-Command docker-compose)) {
    Write-Host "Docker Compose não está instalado. Por favor, instale o Docker Compose primeiro." -ForegroundColor $Red
    exit 1
}

# Parar containers existentes
Write-Host "Parando containers existentes..." -ForegroundColor $Yellow
docker-compose -f docker-compose.prod.yml down

# Rebuildar imagens
Write-Host "Rebuildando imagens..." -ForegroundColor $Yellow
docker-compose -f docker-compose.prod.yml build --no-cache

# Iniciar containers em background
Write-Host "Iniciando containers..." -ForegroundColor $Yellow
docker-compose -f docker-compose.prod.yml up -d

# Aguardar todos os serviços estarem prontos
Write-Host "Aguardando serviços iniciarem..." -ForegroundColor $Yellow
Start-Sleep -Seconds 10

# Verificar status dos containers
Write-Host "Verificando status dos containers..." -ForegroundColor $Yellow
$containerStatus = docker-compose -f docker-compose.prod.yml ps
if ($containerStatus -match "Exit") {
    Write-Host "Alguns containers falharam ao iniciar. Verifique os logs:" -ForegroundColor $Red
    docker-compose -f docker-compose.prod.yml logs
    exit 1
}

Write-Host "Serviços disponíveis em:" -ForegroundColor $Green
Write-Host "Frontend: http://localhost:3000" -ForegroundColor $Green
Write-Host "Backend:  http://localhost:8080" -ForegroundColor $Green
Write-Host "Postgres: localhost:5432" -ForegroundColor $Green 