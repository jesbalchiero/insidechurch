package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"insidechurch/backend/internal/core/domain"
	"insidechurch/backend/internal/handlers"
	"insidechurch/backend/internal/middleware"
	"insidechurch/backend/internal/repositories"
	"insidechurch/backend/internal/routes"
	"insidechurch/backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	// Configurar banco de dados de teste
	dsn := "host=localhost user=postgres password=postgres dbname=insidechurch_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados de teste: %v", err)
	}

	// Limpar tabelas antes dos testes
	db.Migrator().DropTable(&domain.User{})
	db.AutoMigrate(&domain.User{})

	return db
}

func setupRouter(db *gorm.DB) *gin.Engine {
	// Configurar componentes
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	userHandler := handlers.NewUserHandler(userRepo, authService)
	authMiddleware := middleware.NewAuthMiddleware(authService)
	securityMiddleware := middleware.NewSecurityMiddleware()

	// Configurar router
	router := gin.Default()
	routes.SetupRoutes(router, userHandler, authMiddleware, securityMiddleware)

	return router
}

func TestRegisterAndLoginFlow(t *testing.T) {
	// Configurar ambiente de teste
	db := setupTestDB(t)
	router := setupRouter(db)

	// Dados de teste
	registerData := map[string]string{
		"email":    "teste@exemplo.com",
		"password": "senha123",
		"name":     "Usuário Teste",
	}

	// Teste de registro
	t.Run("Registro de usuário", func(t *testing.T) {
		jsonData, _ := json.Marshal(registerData)
		req := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response["token"])
		assert.NotEmpty(t, response["user"])
	})

	// Teste de login
	t.Run("Login de usuário", func(t *testing.T) {
		loginData := map[string]string{
			"email":    registerData["email"],
			"password": registerData["password"],
		}

		jsonData, _ := json.Marshal(loginData)
		req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response["token"])
		assert.NotEmpty(t, response["user"])
	})

	// Teste de acesso protegido
	t.Run("Acesso a rota protegida", func(t *testing.T) {
		// Primeiro fazer login para obter o token
		loginData := map[string]string{
			"email":    registerData["email"],
			"password": registerData["password"],
		}
		jsonData, _ := json.Marshal(loginData)
		req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var loginResponse map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &loginResponse)
		token := loginResponse["token"].(string)

		// Agora testar a rota protegida
		req = httptest.NewRequest("GET", "/api/user", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var userResponse map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &userResponse)
		assert.NoError(t, err)
		assert.NotEmpty(t, userResponse["user"])
	})

	// Teste de login com credenciais inválidas
	t.Run("Login com credenciais inválidas", func(t *testing.T) {
		invalidData := map[string]string{
			"email":    registerData["email"],
			"password": "senha_errada",
		}

		jsonData, _ := json.Marshal(invalidData)
		req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	// Teste de acesso protegido sem token
	t.Run("Acesso sem token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/user", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
