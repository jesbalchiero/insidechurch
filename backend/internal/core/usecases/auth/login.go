package auth

import (
	"errors"
	"os"
	"time"

	"insidechurch/backend/internal/core/domain/entities"
	"insidechurch/backend/internal/core/ports"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("credenciais inválidas")
	ErrInvalidInput       = errors.New("entrada inválida")
)

// LoginUseCase implementa o caso de uso de login
type LoginUseCase struct {
	userRepo ports.UserRepository
}

// NewLoginUseCase cria uma nova instância do caso de uso de login
func NewLoginUseCase(userRepo ports.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		userRepo: userRepo,
	}
}

// Login executa o caso de uso de login
func (uc *LoginUseCase) Login(input LoginInput) (*LoginOutput, error) {
	// Validação de entrada
	if input.Email == "" || input.Password == "" {
		return nil, ErrInvalidInput
	}

	// Buscar usuário pelo email
	user, err := uc.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if user == nil {
		return nil, ErrInvalidCredentials
	}

	// Verificar senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	// Gerar token JWT
	token, err := uc.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		User:  user,
		Token: token,
	}, nil
}

// ValidateToken valida um token JWT
func (uc *LoginUseCase) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return token, nil
}

// generateToken gera um token JWT para o usuário
func (uc *LoginUseCase) generateToken(user *entities.User) (string, error) {
	// Criar claims do token
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
	}

	// Criar token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assinar token
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
