package cache

import (
	"context"
	"time"
)

// Cache define a interface para operações de cache
type Cache interface {
	// Get recupera um valor do cache
	Get(ctx context.Context, key string) (interface{}, error)

	// Set armazena um valor no cache com tempo de expiração
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error

	// Delete remove um valor do cache
	Delete(ctx context.Context, key string) error

	// Exists verifica se uma chave existe no cache
	Exists(ctx context.Context, key string) (bool, error)

	// Flush limpa todo o cache
	Flush(ctx context.Context) error
}
