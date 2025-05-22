package handlers

import (
	"net/http"

	"insidechurch/backend/internal/core/domain"
	"insidechurch/backend/internal/core/interfaces"
	"insidechurch/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo    interfaces.UserRepository
	authService *services.AuthService
}

func NewUserHandler(userRepo interfaces.UserRepository, authService *services.AuthService) *UserHandler {
	return &UserHandler{
		userRepo:    userRepo,
		authService: authService,
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"user"`
}

// Register godoc
// @Summary      Registrar novo usuário
// @Description  Cria um novo usuário na plataforma
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body  RegisterRequest  true  "Dados de registro"
// @Success      201  {object}  RegisterResponse
// @Failure      400  {object}  gin.H
// @Failure      409  {object}  gin.H
// @Router       /api/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	// Validar senha
	if err := h.authService.ValidatePassword(req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se o email já existe
	existingUser, _ := h.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email já cadastrado"})
		return
	}

	// Criar hash da senha
	hashedPassword, err := h.authService.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao processar senha"})
		return
	}

	// Criar usuário
	user := &domain.User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
	}

	if err := h.userRepo.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar usuário"})
		return
	}

	// Gerar token
	token, err := h.authService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar token"})
		return
	}

	// Retornar resposta
	c.JSON(http.StatusCreated, AuthResponse{
		Token: token,
		User: struct {
			ID    uint   `json:"id"`
			Email string `json:"email"`
			Name  string `json:"name"`
		}{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	})
}

// Login godoc
// @Summary      Login do usuário
// @Description  Autentica um usuário e retorna um token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body  LoginRequest  true  "Credenciais de login"
// @Success      200  {object}  LoginResponse
// @Failure      400  {object}  gin.H
// @Failure      401  {object}  gin.H
// @Router       /api/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	// Autenticar usuário
	token, err := h.authService.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciais inválidas"})
		return
	}

	// Buscar usuário para retornar dados
	user, err := h.userRepo.FindByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar usuário"})
		return
	}

	// Retornar resposta
	c.JSON(http.StatusOK, AuthResponse{
		Token: token,
		User: struct {
			ID    uint   `json:"id"`
			Email string `json:"email"`
			Name  string `json:"name"`
		}{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	})
}

// GetUser godoc
// @Summary      Buscar usuário autenticado
// @Description  Retorna os dados do usuário autenticado
// @Tags         user
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  UserResponse
// @Failure      401  {object}  gin.H
// @Router       /api/user [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	// Obter ID do usuário do contexto (definido pelo middleware de autenticação)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "não autorizado"})
		return
	}

	// Buscar usuário
	user, err := h.userRepo.FindByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado"})
		return
	}

	// Retornar resposta
	c.JSON(http.StatusOK, AuthResponse{
		User: struct {
			ID    uint   `json:"id"`
			Email string `json:"email"`
			Name  string `json:"name"`
		}{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	})
}
