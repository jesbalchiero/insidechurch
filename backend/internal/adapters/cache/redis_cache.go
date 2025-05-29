package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache implementa a interface Cache usando Redis
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache cria uma nova instância do RedisCache
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

// Get recupera um valor do cache
func (c *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("erro ao recuperar valor do cache: %w", err)
	}

	var result interface{}
	if err := json.Unmarshal([]byte(val), &result); err != nil {
		return nil, fmt.Errorf("erro ao desserializar valor do cache: %w", err)
	}

	return result, nil
}

// Set armazena um valor no cache
func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("erro ao serializar valor para cache: %w", err)
	}

	if err := c.client.Set(ctx, key, data, expiration).Err(); err != nil {
		return fmt.Errorf("erro ao armazenar valor no cache: %w", err)
	}

	return nil
}

// Delete remove um valor do cache
func (c *RedisCache) Delete(ctx context.Context, key string) error {
	if err := c.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("erro ao remover valor do cache: %w", err)
	}
	return nil
}

// Exists verifica se uma chave existe no cache
func (c *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("erro ao verificar existência da chave no cache: %w", err)
	}
	return exists > 0, nil
}

// Flush limpa todo o cache
func (c *RedisCache) Flush(ctx context.Context) error {
	if err := c.client.FlushAll(ctx).Err(); err != nil {
		return fmt.Errorf("erro ao limpar cache: %w", err)
	}
	return nil
}
