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
	// Buscar usuário pelo email
	user, err := uc.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	if user == nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Verificar senha
	if !uc.checkPasswordHash(input.Password, user.Password) {
		return nil, errors.New("credenciais inválidas")
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
func (uc *LoginUseCase) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token inválido")
}

// checkPasswordHash verifica se a senha corresponde ao hash
func (uc *LoginUseCase) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// generateToken gera um token JWT para o usuário
func (uc *LoginUseCase) generateToken(user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
