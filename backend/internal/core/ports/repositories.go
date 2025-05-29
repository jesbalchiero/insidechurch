package ports

import (
	"insidechurch/backend/internal/core/domain/entities"
)

// UserRepository define a interface para operações de persistência de usuários
type UserRepository interface {
	Create(user *entities.User) error
	FindByID(id uint) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uint) error
}
