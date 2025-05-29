package auth

import (
	"errors"

	"insidechurch/backend/internal/core/domain/entities"
	"insidechurch/backend/internal/core/ports"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUseCase implementa o caso de uso de registro
type RegisterUseCase struct {
	userRepo ports.UserRepository
}

// NewRegisterUseCase cria uma nova instância do caso de uso de registro
func NewRegisterUseCase(userRepo ports.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo: userRepo,
	}
}

// Register executa o caso de uso de registro
func (uc *RegisterUseCase) Register(input RegisterInput) error {
	// Validar senha
	if err := uc.validatePassword(input.Password); err != nil {
		return err
	}

	// Verificar se o email já está em uso
	existingUser, err := uc.userRepo.FindByEmail(input.Email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("email já está em uso")
	}

	// Criar hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Criar usuário
	user := &entities.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	return uc.userRepo.Create(user)
}

// validatePassword valida se a senha atende aos requisitos mínimos
func (uc *RegisterUseCase) validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("a senha deve ter pelo menos 8 caracteres")
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case char == '!' || char == '@' || char == '#' || char == '$' || char == '%' || char == '^' || char == '&' || char == '*':
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return errors.New("a senha deve conter pelo menos uma letra maiúscula, uma minúscula, um número e um caractere especial")
	}

	return nil
}
