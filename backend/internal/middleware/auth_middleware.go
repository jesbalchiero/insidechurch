package middleware

import (
	"net/http"
	"strings"

	"insidechurch/backend/internal/core/usecases/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		// Obter o token do header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token não fornecido"})
			c.Abort()
			return
		}

		// Verificar se o token está no formato correto (Bearer <token>)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "formato de token inválido"})
			c.Abort()
			return
		}

		// Validar o token
		token, err := m.loginUseCase.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			c.Abort()
			return
		}

		// Extrair as claims do token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "claims inválidas"})
			c.Abort()
			return
		}

		// Extrair o ID do usuário das claims
		userID, ok := claims["sub"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ID do usuário inválido"})
			c.Abort()
			return
		}

		// Adicionar o ID do usuário ao contexto
		c.Set("userID", uint(userID))
		c.Next()
	}
}
