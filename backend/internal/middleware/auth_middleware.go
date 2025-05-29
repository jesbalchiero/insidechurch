package middleware

import (
	"net/http"
	"strings"

	"insidechurch/backend/internal/core/usecases/auth"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	loginUseCase *auth.LoginUseCase
}

func NewAuthMiddleware(loginUseCase *auth.LoginUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		loginUseCase: loginUseCase,
	}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Verifica se o header começa com "Bearer "
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		token := parts[1]

		// Valida o token usando o caso de uso de login
		claims, err := m.loginUseCase.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Define o user_id no contexto
		c.Set("user_id", claims["sub"])
		c.Next()
	}
}
