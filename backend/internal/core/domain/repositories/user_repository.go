package repositories

import (
	"insidechurch/backend/internal/core/domain/entities"
)

// UserRepository define as operações de persistência para usuários
type UserRepository interface {
	// Create cria um novo usuário
	Create(user *entities.User) error

	// FindByID busca um usuário pelo ID
	FindByID(id uint) (*entities.User, error)

	// FindByEmail busca um usuário pelo email
	FindByEmail(email string) (*entities.User, error)

	// Update atualiza um usuário existente
	Update(user *entities.User) error

	// Delete remove um usuário
	Delete(id uint) error

	// List retorna uma lista de usuários com paginação
	List(page, limit int) ([]entities.User, int64, error)

	// ExistsByEmail verifica se já existe um usuário com o email informado
	ExistsByEmail(email string) (bool, error)
}
