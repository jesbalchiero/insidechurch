package services

import (
	"errors"
	"os"
	"time"

	"insidechurch/internal/core/domain"
	"insidechurch/internal/core/interfaces"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo interfaces.UserRepository
}

func NewAuthService(userRepo interfaces.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// ValidatePassword verifica se a senha atende aos requisitos mínimos
func (s *AuthService) ValidatePassword(password string) error {
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

// HashPassword gera um hash da senha usando bcrypt
func (s *AuthService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords compara uma senha com seu hash
func (s *AuthService) ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateToken gera um token JWT para o usuário
func (s *AuthService) GenerateToken(user *domain.User) (string, error) {
	// Configurar as claims do token
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
		"iat":   time.Now().Unix(),
	}

	// Criar o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assinar o token com a chave secreta
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken valida um token JWT
func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar se o método de assinatura é o esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

// Authenticate valida as credenciais do usuário e retorna um token JWT
func (s *AuthService) Authenticate(email, password string) (string, error) {
	// Buscar usuário pelo email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	// Verificar a senha
	if err := s.ComparePasswords(user.Password, password); err != nil {
		return "", errors.New("credenciais inválidas")
	}

	// Gerar token JWT
	token, err := s.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
