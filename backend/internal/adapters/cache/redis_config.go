package cache

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// RedisConfig contém as configurações para conexão com o Redis
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// NewRedisConfig cria uma nova configuração do Redis a partir de variáveis de ambiente
func NewRedisConfig() *RedisConfig {
	port, _ := strconv.Atoi(getEnv("REDIS_PORT", "6379"))
	db, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))

	return &RedisConfig{
		Host:     getEnv("REDIS_HOST", "localhost"),
		Port:     port,
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       db,
	}
}

// NewRedisClient cria um novo cliente Redis com as configurações fornecidas
func NewRedisClient(config *RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
}

// getEnv retorna o valor da variável de ambiente ou um valor padrão
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
