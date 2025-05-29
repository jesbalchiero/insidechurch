package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"insidechurch/backend/internal/adapters/handlers"
	"insidechurch/backend/internal/adapters/repositories"
	"insidechurch/backend/internal/core/domain"
	"insidechurch/backend/internal/core/usecases/auth"
	"insidechurch/backend/internal/core/usecases/user"
	"insidechurch/backend/internal/middleware"
	"insidechurch/backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestRouter() *gin.Engine {
	// Configurar banco de dados de teste
	dsn := "host=localhost user=postgres password=postgres dbname=insidechurch_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Limpar tabelas antes dos testes
	db.Exec("DROP TABLE IF EXISTS users CASCADE")
	db.AutoMigrate(&domain.User{})

	// Inicializa o router
	router := gin.Default()

	// Inicializa os repositórios
	userRepo := repositories.NewUserRepository(db)

	// Inicializa os casos de uso
	loginUseCase := auth.NewLoginUseCase(userRepo)
	registerUseCase := auth.NewRegisterUseCase(userRepo)
	getUserUseCase := user.NewGetUserUseCase(userRepo)

	// Inicializa os middlewares
	authMiddleware := middleware.NewAuthMiddleware(loginUseCase)
	securityMiddleware := middleware.NewSecurityMiddleware()

	// Inicializa os handlers
	authHandler := handlers.NewAuthHandler(loginUseCase, registerUseCase)
	userHandler := handlers.NewUserHandler(getUserUseCase)

	// Configura as rotas
	routes.SetupRoutes(router, authHandler, userHandler, authMiddleware, securityMiddleware)

	return router
}

func TestRegisterAndLogin(t *testing.T) {
	router := setupTestRouter()

	// Teste de registro
	registerData := map[string]string{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "Test@123",
	}
	registerBody, _ := json.Marshal(registerData)
	registerReq := httptest.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(registerBody))
	registerReq.Header.Set("Content-Type", "application/json")
	registerW := httptest.NewRecorder()
	router.ServeHTTP(registerW, registerReq)

	assert.Equal(t, http.StatusCreated, registerW.Code)

	// Teste de login
	loginData := map[string]string{
		"email":    "test@example.com",
		"password": "Test@123",
	}
	loginBody, _ := json.Marshal(loginData)
	loginReq := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(loginBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginW := httptest.NewRecorder()
	router.ServeHTTP(loginW, loginReq)

	assert.Equal(t, http.StatusOK, loginW.Code)

	var loginResponse map[string]interface{}
	err := json.Unmarshal(loginW.Body.Bytes(), &loginResponse)
	assert.NoError(t, err)
	assert.NotEmpty(t, loginResponse["token"])

	// Teste de acesso à rota protegida
	userReq := httptest.NewRequest("GET", "/api/user", nil)
	userReq.Header.Set("Authorization", "Bearer "+loginResponse["token"].(string))
	userW := httptest.NewRecorder()
	router.ServeHTTP(userW, userReq)

	assert.Equal(t, http.StatusOK, userW.Code)
}
