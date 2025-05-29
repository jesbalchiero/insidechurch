package handlers

import (
	"net/http"

	"insidechurch/backend/internal/core/usecases/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	loginUseCase    *auth.LoginUseCase
	registerUseCase *auth.RegisterUseCase
}

// NewAuthHandler cria uma nova instância do handler de autenticação
func NewAuthHandler(
	loginUseCase *auth.LoginUseCase,
	registerUseCase *auth.RegisterUseCase,
) *AuthHandler {
	return &AuthHandler{
		loginUseCase:    loginUseCase,
		registerUseCase: registerUseCase,
	}
}

// Login lida com a requisição de login
func (h *AuthHandler) Login(c *gin.Context) {
	var input auth.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	output, err := h.loginUseCase.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  output.User,
		"token": output.Token,
	})
}

// Register lida com a requisição de registro
func (h *AuthHandler) Register(c *gin.Context) {
	var input auth.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	if err := h.registerUseCase.Register(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "usuário criado com sucesso"})
}
