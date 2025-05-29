package auth

import (
	"insidechurch/backend/internal/core/domain/entities"
)

// LoginInput representa os dados necessários para o login
type LoginInput struct {
	Email    string
	Password string
}

// LoginOutput representa o resultado do login
type LoginOutput struct {
	User  *entities.User
	Token string
}

// RegisterInput representa os dados necessários para o registro
type RegisterInput struct {
	Name     string
	Email    string
	Password string
}

// AuthUseCase define a interface para os casos de uso de autenticação
type AuthUseCase interface {
	Login(input LoginInput) (*LoginOutput, error)
	Register(input RegisterInput) error
	ValidateToken(token string) (*entities.User, error)
}
